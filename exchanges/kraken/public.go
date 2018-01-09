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
		Result map[string]interface{} `json:"result"`
	}
)

var marketList []crypto.Market

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []crypto.Market{}, errors.Wrap(err, "couldn't fetch markets")
	}

	defer res.Body.Close()

	var r marketsResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return []crypto.Market{}, errors.Wrap(err, "unable to json-decode response")
	}

	for symbol := range r.Result {
		market, err := crypto.NewMarketFromSymbol(symbol)
		if err != nil {
			log.Debugln("unable to parse symbol:", err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

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
		return ob, errors.Wrap(err, "unable to GET orderbook")
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return ob, errors.Wrap(err, "unable to json-decode response")
	}

	prefAsset, err := nonsensePrefix(assetSymbol)
	if err != nil {
		return ob, errors.Wrap(err, "cannot prefix asset")
	}

	prefPrice, err := nonsensePrefix(priceSymbol)
	if err != nil {
		return ob, errors.Wrap(err, "cannot prefix pricedIn")
	}

	marketKey := fmt.Sprintf("%s%s", prefAsset, prefPrice)
	for market, orderResp := range r.Result {
		if len(r.Result) > 1 {
			log.Debugf("more than one market returned for %s…", m)
		}

		if market == marketKey || len(r.Result) == 1 {
			if market != marketKey {
				log.Debugf("non-standard market symbol returned: %s instead of %s", market, marketKey)
			}

			ob, err = orderbook.Normalise(orderResp.Asks, orderResp.Bids)
			if err != nil {
				err = errors.Wrap(err, "unable to normalise Order Book")
			}

			return
		}
	}

	return ob, errors.Errorf("%s market not found.", marketKey)
}
