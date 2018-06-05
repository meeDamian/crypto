package livecoin

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

const balancesUrl = "https://api.livecoin.net/payment/balances"

type (
	errorResp struct {
		Success bool    `json:"success"`
		Error   *string `json:"error"`
	}

	balancesResp []struct {
		Type     string  `json:"type"`
		Currency string  `json:"currency"`
		Value    float64 `json:"value"`
	}
)

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "GET", balancesUrl, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		var r errorResp
		err = json.NewDecoder(res.Body).Decode(&r)
		if err != nil {
			return
		}

		return balances, errors.New(*r.Error)
	}

	var bals balancesResp
	err = json.NewDecoder(res.Body).Decode(&bals)
	if err != nil {
		return
	}

	var tmpBalances = make(crypto.Balances)
	for _, b := range bals {
		if b.Type == "trade" || b.Type == "available_withdrawal" {
			continue
		}

		balance, ok := tmpBalances[b.Currency]
		if !ok {
			balance = crypto.Balance{}
		}

		if b.Type == "total" {
			balance.Total = b.Value
		}

		if b.Type == "available" {
			balance.Available = b.Value
		}

		tmpBalances[b.Currency] = balance
	}

	balances = make(crypto.Balances)
	for name, bb := range tmpBalances {
		err := balances.Add(name, bb.Available, bb.Total, nil)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", name, err)
		}
	}

	return
}
