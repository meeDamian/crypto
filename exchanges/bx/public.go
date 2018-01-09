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

func (p pairings) getId(market crypto.Market) (int, error) {
	for id, m := range p {
		if m.Asset == market.Asset && m.PricedIn == market.PricedIn {
			return id, nil
		}
	}

	return 0, errors.Errorf("requested market(%s%s) not found", market.Asset, market.PricedIn)
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	if len(marketPairings) == 0 {
		return ob, errors.New("call to bx.OrderBook() requires prior call to bx.Markets()")
	}

	id, err := marketPairings.getId(m)
	if err != nil {
		return ob, errors.Wrapf(err, "unable to get %s market pairing", Domain)
	}

	url := fmt.Sprintf(orderBookUrl, id)
	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrap(err, "unable to fetch Order Book")
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

	var ms map[string]pairing
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	marketPairings = make(map[int]crypto.Market)

	for _, m := range ms {
		market := crypto.NewMarket(m.Asset, m.PricedIn)

		marketPairings[m.Id] = market
		marketList = append(marketList, market)
	}

	return marketList, nil
}
