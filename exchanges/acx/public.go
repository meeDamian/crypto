package acx

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const (
	marketsUrl   = "https://acx.io/api/v2/markets.json"
	orderBookUrl = "https://acx.io/api/v2/depth.json?market=%s%s"
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
		marketList, err = crypto.AppendMarket(marketList, m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
		}
	}

	return marketList, nil
}

func morph(currency string) string {
	return strings.ToLower(currency)
}

func OrderBook(m crypto.Market) (orderbook.OrderBook, error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), morph(m.PricedIn))
	return orderbook.Download(url)
}
