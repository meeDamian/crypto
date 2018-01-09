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
	orderBookUrl = "https://api.liqui.io/api/3/depth/%s_%s"
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
		market := crypto.NewMarket(symbols[0], symbols[1])

		if d.Hidden == 1 {
			log.Debugf("Skipping hidden market: %s/%s", market.Asset, market.PricedIn)
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
	asset, price := morph(m.Asset), morph(m.PricedIn)
	url := fmt.Sprintf(orderBookUrl, asset, price)

	res, err := utils.NetClient().Get(url)
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
		if name == fmt.Sprintf("%s_%s", asset, price) {
			return orderbook.Normalise(ob.Asks, ob.Bids)
		}
	}

	return ob, errors.New("no matching market found in response…")
}
