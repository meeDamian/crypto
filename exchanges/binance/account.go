package binance

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

const balancesUrl = "https://api.binance.com/api/v3/account"

type accountInfo struct {
	Balances []struct {
		Currency  string `json:"asset"`
		Available string `json:"free"`
		Locked    string `json:"locked"`
	} `json:"balances"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "GET", balancesUrl, nil)
	if err != nil {
		return balances, err
	}

	defer res.Body.Close()

	var r accountInfo
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return balances, errors.Wrap(err, "can't json-decode response")
	}

	balances = make(crypto.Balances)
	for _, b := range r.Balances {
		err := balances.Add(b.Currency, b.Available, nil, b.Locked)
		if err != nil {
			log.Debugf("skipping balance of %s, due to: %v", b.Currency, err)
		}
	}
	return
}
