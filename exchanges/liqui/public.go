package liqui

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

type (
	info struct {
		Pairs map[string]struct {
			Hidden int `json:"hidden"`
		} `json:"pairs"`
	}

	obResp map[string]orderbook.ObResponse
)

const (
	infoUrl      = "https://api.liqui.io/api/3/info"
	orderBookUrl = "https://api.liqui.io/api/3/depth/%s"
)

var marketList []crypto.Market

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(infoUrl)
	if err != nil {
		return []crypto.Market{}, err
	}

	defer res.Body.Close()

	var ts info
	err = json.NewDecoder(res.Body).Decode(&ts)
	if err != nil {
		return
	}

	for pair, d := range ts.Pairs {
		symbols := strings.Split(pair, "_")
		market, err := crypto.NewMarket(symbols[0], symbols[1])
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", market.Asset, market.PricedIn, err)
			continue
		}

		if d.Hidden == 1 {
			log.Debugf("skipping market %s/%s: marked as hidden by exchange", market.Asset, market.PricedIn)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

func morph(name string) string {
	return strings.ToLower(currencies.Morph(name, aliases))
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	symbol := fmt.Sprintf("%s_%s", morph(m.Asset), morph(m.PricedIn))

	res, err := utils.NetClient().Get(fmt.Sprintf(orderBookUrl, symbol))
	if err != nil {
		return ob, err
	}

	defer res.Body.Close()

	var r obResp
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return ob, err
	}

	for name, ob := range r {
		if name == symbol {
			return orderbook.Normalise(ob.Asks, ob.Bids)
		}
	}

	return ob, errors.Errorf("no matching market(%s) found in response", symbol)
}
