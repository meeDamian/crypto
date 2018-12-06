package yobit

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
	marketsUrl   = "https://yobit.net/api/3/info"
	orderBookUrl = "https://yobit.net/api/3/depth/%s"
)

type (
	marketRes struct {
		Pairs map[string]interface{} `json:"pairs"`
	}

	obRes map[string]orderbook.ObResponse
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

	var ms marketRes
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for symbol := range ms.Pairs {
		market, err := markets.NewFromSymbol(symbol)
		if err != nil {
			log.Debugf("skipping symbol %s: %v", symbol, err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

func morph(name string) string {
	return currencies.Morph(name, aliases)
}

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	symbol := strings.ToLower(fmt.Sprintf("%s_%s", morph(m.Asset), morph(m.PricedIn)))

	url := fmt.Sprintf(orderBookUrl, symbol)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		return orderbook.OrderBook{}, err
	}

	defer res.Body.Close()

	var r obRes
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return orderbook.OrderBook{}, err
	}

	if val, ok := r[symbol]; ok {
		return orderbook.Normalise(val.Asks, val.Bids)
	}

	return orderbook.OrderBook{}, errors.Errorf("symbol %s not available in response", symbol)
}
