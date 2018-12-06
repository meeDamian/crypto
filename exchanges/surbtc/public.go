package surbtc

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"

	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	orderBookUrl = "https://www.surbtc.com/api/v2/markets/%s-%s/order_book.json"
	marketsUrl   = "https://www.surbtc.com/api/v2/markets.json"
)

type (
	marketResponse struct {
		Markets []struct {
			Asset    string `json:"base_currency"`
			PricedIn string `json:"quote_currency"`
		} `json:"markets"`
	}

	obResponse struct {
		OrderBook orderbook.ObResponse `json:"order_book"`
	}
)

var marketList []markets.Market

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, m.Asset, m.PricedIn)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		return
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	return orderbook.Normalise(r.OrderBook.Asks, r.OrderBook.Bids)
}

func Markets() (_ []markets.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []markets.Market{}, err
	}

	defer res.Body.Close()

	var ms marketResponse
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms.Markets {
		marketList, err = markets.Append(marketList, m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
		}
	}

	return marketList, nil
}
