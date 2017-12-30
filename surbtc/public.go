package surbtc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
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

var marketList []crypto.Market

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, m.Asset, m.PricedIn)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		err = errors.Wrap(err, "unable to GET orderbook")
		return
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		err = errors.Wrap(err, "unable to decode response")
		return
	}

	ob, err = orderbook.Normalise(r.OrderBook.Asks, r.OrderBook.Bids)
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

		var ms marketResponse
		err = json.NewDecoder(res.Body).Decode(&ms)
		if err != nil {
			return
		}

		for _, m := range ms.Markets {
			marketList = append(marketList, crypto.NewMarket(m.Asset, m.PricedIn))
		}
	}

	return marketList, nil
}
