package bitstamp

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	Pair string `json:"name"`
}

var marketList []crypto.Market

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, strings.ToLower(m.Asset), strings.ToLower(m.PricedIn))

	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrapf(err, "unable to fetch %s Order Book", Domain)
	}

	return
}

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) == 0 {
		var res *http.Response
		res, err = utils.NetClient().Get(marketsUrl)
		if err != nil {
			return
		}

		defer res.Body.Close()

		var ms []market
		err = json.NewDecoder(res.Body).Decode(&ms)
		if err != nil {
			return
		}

		for _, m := range ms {
			pair := strings.Split(m.Pair, "/")

			marketList = append(marketList, crypto.NewMarket(pair[0], pair[1]))
		}
	}

	return marketList, nil
}
