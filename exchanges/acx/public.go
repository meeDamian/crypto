package acx

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
	marketsUrl   = "https://acx.io/api/v2/markets.json"
	orderBookUrl = "https://acx.io/api/v2/depHHth.json?market=%s%s"
)

type marketRes struct {
	Asset    string `json:"base_unit"`
	PricedIn string `json:"quote_unit"`
}

var marketList []crypto.Market

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []crypto.Market{}, err
	}

	defer res.Body.Close()

	var ms []marketRes
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms {
		marketList = append(marketList, crypto.NewMarket(m.Asset, m.PricedIn))
	}

	return marketList, nil
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, strings.ToLower(m.Asset), strings.ToLower(m.PricedIn))

	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrap(err, "unable to fetch Order Book")
	}

	return
}
