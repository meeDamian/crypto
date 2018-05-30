package livecoin

import (
	"encoding/json"
	"fmt"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	infoUrl          = "https://api.livecoin.net/exchange/ticker"
	orderBookUrl     = "https://api.livecoin.net/exchange/order_book?currencyPair=%s/%s"
	allOrderBooksUrl = "https://api.livecoin.net/exchange/all/order_book"
)

type (
	markets []struct {
		Symbol string `json:"symbol"`
	}

	allObResp map[string]orderbook.ObResponse
)

var marketList []crypto.Market

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(infoUrl)
	if err != nil {
		return []crypto.Market{}, err
	}

	defer res.Body.Close()

	var ms markets
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms {
		market, err := crypto.NewMarketFromSymbol(m.Symbol)
		if err != nil {
			log.Debugf("skipping symbol %s: %v", m.Symbol, err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

func OrderBook(m crypto.Market) (orderbook.OrderBook, error) {
	url := fmt.Sprintf(orderBookUrl, m.Asset, m.PricedIn)
	return orderbook.Download(url)
}

// WARNING: returns 10 top orders per OB side MAX
func AllOrderBooks() (obs map[crypto.Market]orderbook.OrderBook, err error) {
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

	obs = make(map[crypto.Market]orderbook.OrderBook)
	for pair, rawOb := range r {
		market, err := crypto.NewMarketFromSymbol(pair)
		if err != nil {
			log.Errorf("can't process symbol %s into a market: %v", pair, err)
			continue
		}

		obs[market], err = orderbook.Normalise(rawOb.Asks, rawOb.Bids)
		if err != nil {
			log.Errorf("can't process %s orderbook: %v", market, err)
			continue
		}
	}

	return
}
