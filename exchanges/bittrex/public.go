package bittrex

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

const (
	orderBookUrl = "https://bittrex.com/api/v1.1/public/getorderbook?market=%s-%s&type=both"
	marketsUrl   = "https://bittrex.com/api/v1.1/public/getmarkets"
)

type (
	market struct {
		Asset    string `json:"MarketCurrency"`
		PricedIn string `json:"BaseCurrency"`
		IsActive bool   `json:"IsActive"`
	}

	marketResp struct {
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Result  []market `json:"result"`
	}

	result struct {
		Asks []interface{} `json:"sell"`
		Bids []interface{} `json:"buy"`
	}

	obResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Result  result `json:"result"`
	}
)

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

	var ms marketResp
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	if !ms.Success {
		return []crypto.Market{}, errors.Errorf("can't fetch markets: %s", ms.Message)
	}

	for _, m := range ms.Result {
		if !m.IsActive {
			log.Debugf("skipping market %s/%s: marked as not active by exchange", m.Asset, m.PricedIn)
			continue
		}

		marketList, err = crypto.AppendMarket(marketList, m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
		}
	}

	return marketList, nil
}

func morph(name string) string {
	return strings.ToLower(currencies.Morph(name, aliases))
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.PricedIn), morph(m.Asset))

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

	return orderbook.Normalise(r.Result.Asks, r.Result.Bids)
}
