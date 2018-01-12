package bitstamp

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	orderBookUrl = "https://www.bitstamp.net/api/v2/order_book/%s%s/"
	marketsUrl   = "https://www.bitstamp.net/api/v2/trading-pairs-info/"
)

type market struct {
	Symbol string `json:"name"`
}

var marketList []crypto.Market

func morph(name string) string {
	return strings.ToLower(name)
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), morph(m.PricedIn))
	return orderbook.Download(url)
}

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []crypto.Market{}, err
	}

	defer res.Body.Close()

	var ms []market
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
