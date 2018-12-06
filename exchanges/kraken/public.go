package kraken

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"
	"strings"

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
		Error  *[]string              `json:"error"`
	}
)

var marketList []markets.Market

func Markets() (_ []markets.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []markets.Market{}, err
	}

	defer res.Body.Close()

	var r marketsResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	if r.Error != nil && len(*r.Error) > 0 {
		return marketList, errors.New(strings.Join(*r.Error, ", "))
	}

	for symbol := range r.Result {
		market, err := markets.NewFromSymbol(symbol)
		if err != nil {
			log.Debugf("skipping symbol %s: %v", symbol, err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

func addPrefix(name string) (string, error) {
	c, err := currencies.Get(name)
	if err != nil {
		return "", errors.Wrap(err, "unable to add Z/X prefix")
	}

	prefix := "X"
	if c.IsFiat {
		prefix = "Z"
	}

	return fmt.Sprintf("%s%s", prefix, name), nil
}

func removePrefix(prefixedName string) (string, error) {
	ret := prefixedName
	if len(ret) == 4 && (ret[0] == 'Z' || ret[0] == 'X') {
		ret = ret[1:]
	}

	c, err := currencies.Get(ret)
	if err != nil {
		return "", errors.Wrapf(err, "can't extract valid currency from %s", prefixedName)
	}

	return c.Name, nil
}

func morph(name string) string {
	return currencies.Morph(name, aliases)
}

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	assetSymbol, priceSymbol := morph(m.Asset), morph(m.PricedIn)
	url := fmt.Sprintf(orderBookUrl, assetSymbol, priceSymbol)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		return ob, err
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	prefAsset, err := addPrefix(assetSymbol)
	if err != nil {
		return ob, errors.Wrap(err, "can't prefix asset")
	}

	prefPrice, err := addPrefix(priceSymbol)
	if err != nil {
		return ob, errors.Wrap(err, "can't prefix price")
	}

	if len(r.Result) > 1 {
		log.Debugf("more than one market returned for %s", m)
	}

	marketKey := fmt.Sprintf("%s%s", prefAsset, prefPrice)
	for market, orderResp := range r.Result {
		if market == marketKey || len(r.Result) == 1 {
			if market != marketKey {
				log.Debugf("non-standard market symbol returned: %s instead of %s", market, marketKey)
			}

			return orderbook.Normalise(orderResp.Asks, orderResp.Bids)
		}
	}

	return ob, errors.Errorf("can't find %s in response", marketKey)
}
