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
	marketList     []crypto.Market
	marketPairings []market
)

func getId(market crypto.Market) (string, error) {
	for _, m := range marketPairings {
		if m.Asset == market.Asset && m.PricedIn == market.PricedIn {
			return m.Id, nil
		}
	}

	return "", errors.Errorf("pairing for requested market(%s) not found", market)
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

	return orderbook.Normalise(r.Asks, r.Bids)
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
		market, err := crypto.NewMarketWithError(m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
			continue
		}

		// remember pairings for disabled markets, but…
		marketPairings = append(marketPairings, m)

		if m.Disabled {
			log.Debugf("skipping market %s/%s: marked as disabled by exchange", m.Asset, m.PricedIn)
			continue
		}

		// …don't return disabled markets
		marketList = append(marketList, market)
	}

	return marketList, nil
}
