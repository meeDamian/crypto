package livecoin

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"

	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	infoUrl          = "https://api.livecoin.net/exchange/ticker"
	orderBookUrl     = "https://api.livecoin.net/exchange/order_book?currencyPair=%s/%s"
	allOrderBooksUrl = "https://api.livecoin.net/exchange/all/order_book"
)

type (
	markets2 []struct {
		Symbol string `json:"symbol"`
	}

	allObResp map[string]orderbook.ObResponse
)

var marketList []markets.Market

func Markets() (_ []markets.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(infoUrl)
	if err != nil {
		return []markets.Market{}, err
	}

	defer res.Body.Close()

	var ms markets2
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms {
		market, err := markets.NewFromSymbol(m.Symbol)
		if err != nil {
			log.Debugf("skipping symbol %s: %v", m.Symbol, err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

func morph(name string) string {
	return currencies.Morph(name, aliases)
}

func OrderBook(m markets.Market) (orderbook.OrderBook, error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), morph(m.PricedIn))
	return orderbook.Download(url)
}

func AllOrderBooks() (obs map[markets.Market]orderbook.OrderBook, err error) {
	log.Warningf("WARNING: %s.AllOrderBooks() returns at most %d top orders on each side of the Order Book.", Domain, 10)

	res, err := utils.NetClient().Get(allOrderBooksUrl)
	if err != nil {
		return obs, err
	}

	defer res.Body.Close()

	var r allObResp
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	obs = make(map[markets.Market]orderbook.OrderBook)
	for pair, rawOb := range r {
		market, err := markets.NewFromSymbol(pair)
		if err != nil {
			log.Debugf("can't process symbol %s into a market: %v", pair, err)
			continue
		}

		obs[market], err = orderbook.Normalise(rawOb.Asks, rawOb.Bids)
		if err != nil {
			log.Debugf("can't process %s orderbook: %v", market, err)
			continue
		}
	}

	return
}
