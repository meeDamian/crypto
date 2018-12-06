package bitfinex

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"
	"strings"

	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	marketsUrl   = "https://api.bitfinex.com/v1/symbols"
	orderBookUrl = "https://api.bitfinex.com/v1/book/%s%s?group=1&limit_bids=100&limit_asks=100"
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

	var ms []string
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms {
		market, err := markets.NewFromSymbol(strings.ToUpper(m))
		if err != nil {
			log.Debugf("skipping symbol %s: %v", m, err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

func morph(name string) string {
	return currencies.Morph(name, aliases)
}

// TODO: requires rate-limiting handled
func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), morph(m.PricedIn))
	return orderbook.Download(url)
}
