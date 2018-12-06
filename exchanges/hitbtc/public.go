package hitbtc

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"

	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	marketsUrl   = "https://api.hitbtc.com/api/2/public/symbol"
	orderBookUrl = "https://api.hitbtc.com/api/2/public/orderbook/%s%s"
)

type (
	market struct {
		Asset    string `json:"baseCurrency"`
		PricedIn string `json:"quoteCurrency"`
	}

	obResp struct {
		Asks []interface{} `json:"ask"`
		Bids []interface{} `json:"bid"`
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

	var ms []market
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms {
		if m.Asset == currencies.Bcc {
			log.Debugf("skipping BitConnect market %s/%s…", m.Asset, m.PricedIn)
			continue
		}

		marketList, err = markets.Append(marketList, m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
		}
	}

	return marketList, nil
}

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, m.Asset, m.PricedIn)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		return ob, err
	}

	defer res.Body.Close()

	var r obResp
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	return orderbook.Normalise(r.Asks, r.Bids)
}
