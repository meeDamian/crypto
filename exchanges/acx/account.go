package acx

import (
	"encoding/json"
	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

type (
	account struct {
		Currency string `json:"currency"`
		Balance  string `json:"balance"`
		Locked   string `json:"locked"`
	}

	me struct {
		Sn        string    `json:"sn"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Activated bool      `json:"activated"`
		Accounts  []account `json:"accounts"`
	}
)

const meUrl = "https://acx.io/api/v2/members/me.json"

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "GET", meUrl, nil)
	if err != nil {
		return balances, err
	}

	defer res.Body.Close()

	var m me
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return balances, errors.Wrap(err, "can't json-decode response")
	}

	balances = make(crypto.Balances)
	for _, b := range m.Accounts {
		err := balances.Add(b.Currency, b.Balance, nil, b.Locked)
		if err != nil {
			log.Debugf("skipping balance of %s, due to: %v", b.Currency, err)
		}
	}

	return
}
