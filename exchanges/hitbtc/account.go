package hitbtc

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
)

type balance struct {
	Currency  string `json:"currency"`
	Available string `json:"available"`
	Locked    string `json:"reserved"`
}

const balanceUrl = "https://api.hitbtc.com/api/2/account/balance"

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "GET", balanceUrl, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var bs []balance
	err = json.NewDecoder(res.Body).Decode(&bs)
	if err != nil {
		return
	}

	balances = make(crypto.Balances)
	for _, b := range bs {
		err := balances.Add(b.Currency, b.Available, nil, b.Locked)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", b.Currency, err)
		}
	}

	return
}
