package quoinex

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const (
	marketsUrl   = "https://api.quoine.com/products"
	orderBookUrl = "https://api.quoine.com/products/%s/price_levels?full=1"
)

type (
	market struct {
		Id       string `json:"id"`
		Asset    string `json:"base_currency"`
		PricedIn string `json:"quoted_currency"`
		Disabled bool   `json:"disabled"`
	}

	obResp struct {
		Asks []interface{} `json:"sell_price_levels"`
		Bids []interface{} `json:"buy_price_levels"`
	}
)

var (
	marketList []crypto.Market
	pairings   []market
)

func getId(market crypto.Market) (string, error) {
	for _, m := range pairings {
		if m.Asset == market.Asset && m.PricedIn == market.PricedIn {
			return m.Id, nil
		}
	}

	return "", errors.Errorf("requested market(%s) not found", market)
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	id, err := getId(m)
	if err != nil {
		return
	}

	url := fmt.Sprintf(orderBookUrl, id)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		return
	}

	defer res.Body.Close()

	var r obResp
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	ob, err = orderbook.Normalise(r.Asks, r.Bids)
	if err != nil {
		err = errors.Wrapf(err, "unable to fetch %s Order Book", Domain)
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

	var markets []market
	err = json.NewDecoder(res.Body).Decode(&markets)
	if err != nil {
		return
	}

	for _, m := range markets {
		pairings = append(pairings, m)

		if m.Disabled {
			log.Debugf("skipping disabled market %s/%s", m.Asset, m.PricedIn)
			continue
		}

		marketList = append(marketList, crypto.NewMarket(m.Asset, m.PricedIn))
	}

	return marketList, nil
}
