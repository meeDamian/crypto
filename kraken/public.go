package kraken

import (
	"encoding/json"
	"fmt"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const (
	orderBookUrl = "https://api.kraken.com/0/public/Depth?pair=%s%s"
	marketsUrl   = "https://api.kraken.com/0/public/AssetPairs"
)

type (
	obResponse struct {
		Result map[string]orderbook.ObResponse `json:"result"`
	}

	marketsResponse struct {
		Result map[string]struct {
			Pair  string `json:"altname"`
			Asset string `json:"base"`
			Price string `json:"quote"`
		} `json:"result"`
	}
)

var aliases = []string{currencies.Xbt}

func nonsensePrefix(curr string) (string, error) {
	c, err := currencies.Get(curr)
	if err != nil {
		return "", errors.Wrap(err, "unable to add kraken prefix")
	}

	prefix := "X"
	if c.IsFiat {
		prefix = "Z"
	}

	return fmt.Sprintf("%s%s", prefix, curr), nil
}

func removeKrakenNonsense(val string) (string, error) {
	ret := val
	if len(ret) == 4 && (ret[0] == 'Z' || ret[0] == 'X') {
		ret = ret[1:]
	}

	c, err := currencies.Get(ret)
	if err != nil {
		return "", errors.Wrapf(err, "couldn't extract valid currency from %s", val)
	}

	return c.Name, nil
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	assetSymbol := currencies.Morph(m.Asset, aliases)
	priceSymbol := currencies.Morph(m.PricedIn, aliases)

	url := fmt.Sprintf(orderBookUrl, assetSymbol, priceSymbol)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		err = errors.Wrap(err, "unable to GET orderbook")
		return
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		err = errors.Wrap(err, "unable to decode response")
		return
	}

	prefAsset, err := nonsensePrefix(assetSymbol)
	if err != nil {
		err = errors.Wrap(err, "cannot prefix asset")
		return
	}

	prefPrice, err := nonsensePrefix(priceSymbol)
	if err != nil {
		err = errors.Wrap(err, "cannot prefix pricedIn")
		return
	}

	marketKey := fmt.Sprintf("%s%s", prefAsset, prefPrice)
	for market, orderResp := range r.Result {
		if len(r.Result) > 1 {
			crypto.Log().Debugln("more than one Kraken market returned…")
		}

		if market == marketKey || len(r.Result) == 1 {
			if market != marketKey {
				crypto.Log().Debugf("non-standard Kraken market key returned: %s instead of %s", market, marketKey)
			}

			ob, err = orderbook.Normalise(orderResp.Asks, orderResp.Bids)
			if err != nil {
				err = errors.Wrap(err, "unable to normalise Kraken's OrderBook")
			}

			return
		}
	}

	err = errors.Errorf("%s market not found.", marketKey)

	return
}

func Markets() (ms []crypto.Market, err error) {
	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		err = errors.Wrap(err, "couldn't fetch markets")
		return
	}

	defer res.Body.Close()

	var r marketsResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		err = errors.Wrap(err, "unable to decode response")
		return
	}

	pairs := make(map[string]bool)
	for _, p := range r.Result {
		asset, err := removeKrakenNonsense(p.Asset)
		if err != nil {
			continue
		}

		price, err := removeKrakenNonsense(p.Price)
		if err != nil {
			continue
		}

		// skip adding if already added
		_, ok := pairs[asset+price]
		if ok {
			continue
		}
		pairs[asset+price] = true

		ms = append(ms, crypto.NewMarket(asset, price))
	}

	return
}
