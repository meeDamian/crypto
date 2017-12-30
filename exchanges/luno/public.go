package luno

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

type tickers struct {
	Tickers []struct {
		Pair string `json:"pair"`
	} `json:"tickers"`
}

const (
	// yes, API is on a different url than the website ¯\_(ツ)_/¯
	orderBookUrl = "https://api.mybitx.com/api/1/orderbook?pair=%s%s"
	marketsUrl   = "https://api.mybitx.com/api/1/tickers"
)

var (
	marketList []crypto.Market
	aliases    = []string{currencies.Xbt}
)

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl,
		currencies.Morph(m.Asset, aliases),
		currencies.Morph(m.PricedIn, aliases),
	)

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

		var ts tickers
		err = json.NewDecoder(res.Body).Decode(&ts)
		if err != nil {
			return
		}

		for _, m := range ts.Tickers {
			marketList = append(marketList, crypto.NewMarket(m.Pair[:3], m.Pair[3:]))
		}
	}

	return marketList, nil
}
