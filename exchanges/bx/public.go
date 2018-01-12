package bx

import (
	"encoding/json"
	"fmt"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const (
	orderBookUrl = "https://bx.in.th/api/orderbook/?pairing=%d"
	marketsUrl   = "https://bx.in.th/api/pairing/"
)

type (
	pairing struct {
		Id       int    `json:"pairing_id"`
		Asset    string `json:"secondary_currency"`
		PricedIn string `json:"primary_currency"`
	}

	pairings map[int]crypto.Market
)

var (
	marketList     []crypto.Market
	marketPairings pairings
)

func getId(market crypto.Market) (int, error) {
	for id, m := range marketPairings {
		if m.Asset == market.Asset && m.PricedIn == market.PricedIn {
			return id, nil
		}
	}

	return 0, errors.Errorf("pairing for requested market(%s) not found", market)
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	id, err := getId(m)
	if err != nil {
		return ob, err
	}

	return orderbook.Download(fmt.Sprintf(orderBookUrl, id))
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

	var ms map[string]pairing
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	marketPairings = make(map[int]crypto.Market)
	for _, m := range ms {
		market, err := crypto.NewMarketWithError(m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
			continue
		}

		marketPairings[m.Id] = market
		marketList = append(marketList, market)
	}

	return marketList, nil
}
