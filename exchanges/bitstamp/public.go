package bitstamp

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const (
	orderBookUrl = "https://www.bitstamp.net/api/v2/order_book/%s%s/"
	marketsUrl   = "https://www.bitstamp.net/api/v2/trading-pairs-info/"
)

type market struct {
	Symbol string `json:"name"`
}

var marketList []crypto.Market

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, strings.ToLower(m.Asset), strings.ToLower(m.PricedIn))

	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrap(err, "unable to fetch Order Book")
	}

	return
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
			log.Debugln("unable to parse symbol:", err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}
