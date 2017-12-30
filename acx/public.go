package acx

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
	orderBookUrl = "https://acx.io/api/v2/depth.json?market=%s%s"
	marketsUrl   = "https://acx.io/api/v2/markets.json"
)

type market struct {
	Asset    string `json:"base_unit"`
	PricedIn string `json:"quote_unit"`
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
