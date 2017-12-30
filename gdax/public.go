package gdax

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
	orderBookUrl = "https://api.gdax.com/products/%s-%s/book?level=2"
	marketsUrl   = "https://api.gdax.com/products"
)

type market struct {
	Asset    string `json:"base_currency"`
	PricedIn string `json:"quote_currency"`
}

var marketList []crypto.Market

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
			marketList = append(marketList, crypto.NewMarket(m.Asset, m.PricedIn))
		}
	}

	return marketList, nil
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, strings.ToLower(m.Asset), strings.ToLower(m.PricedIn))

	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrapf(err, "unable to fetch %s Order Book", Domain)
	}

	return
}
