package bx

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"

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
		Id       string `json:"pairing_id"`
		Asset    string `json:"secondary_currency"`
		PricedIn string `json:"primary_currency"`
		Active   bool   `json:"active"`
	}

	pairings map[string]markets.Market
)

var (
	marketList     []markets.Market
	marketPairings pairings
)

func getId(market markets.Market) (string, error) {
	for id, m := range marketPairings {
		if m.Asset == market.Asset && m.PricedIn == market.PricedIn {
			return id, nil
		}
	}

	return "", errors.Errorf("pairing for requested market(%s) not found", market)
}

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	id, err := getId(m)
	if err != nil {
		return ob, err
	}

	return orderbook.Download(fmt.Sprintf(orderBookUrl, id))
}

func Markets() (_ []markets.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []markets.Market{}, err
	}

	defer res.Body.Close()

	var ms map[string]pairing
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	_, ok := ms["success"]
	if ok {
		return []markets.Market{}, errors.Errorf("market DL rate limited")
	}

	marketPairings = make(pairings)
	for _, m := range ms {
		market, err := markets.New(m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
			continue
		}

		marketPairings[m.Id] = market
		if !m.Active {
			log.Debugf("skipping market %s/%s: marked as hidden by exchange", m.Asset, m.PricedIn)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}
