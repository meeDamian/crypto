package liqui

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	obResp map[string]interface{}
)

const (
	infoUrl      = "https://api.liqui.io/api/3/info"
	orderBookUrl = "https://api.liqui.io/api/3/depth/%s"

	errorTooManyRequests = "Requests too often"
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
		market, err := crypto.NewMarketFromSymbol(pair)
		if err != nil {
			log.Debugf("skipping market %s: %v", pair, err)
			continue
		}

		if d.Hidden == 1 {
			log.Debugf("skipping market %s: marked as hidden by exchange",market)
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

	if errText, ok := r["error"].(string); ok {
		if errText == errorTooManyRequests {
			return ob, errors.New(http.StatusText(http.StatusTooManyRequests))
		}

		return ob, errors.New(errText)
	}

	for name, ob := range r {
		if name == symbol {
			obSides := ob.(map[string]interface{})

			return orderbook.Normalise(
				obSides["asks"].([]interface{}),
				obSides["bids"].([]interface{}),
			)
		}
	}

	return ob, errors.Errorf("no matching market(%s) found in response", symbol)
}
