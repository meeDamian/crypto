package gdax

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"
	"strings"

	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	orderBookUrl = "https://api.gdax.com/products/%s-%s/book?level=2"
	marketsUrl   = "https://api.gdax.com/products"
)

type market struct {
	Asset    string `json:"base_currency"`
	PricedIn string `json:"quote_currency"`
}

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
		marketList, err = markets.Append(marketList, m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
		}
	}

	return marketList, nil
}

func morph(name string) string {
	return strings.ToLower(name)
}

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), morph(m.PricedIn))
	return orderbook.Download(url)
}
