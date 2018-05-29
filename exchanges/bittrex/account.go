package bittrex

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
)

const balancesUrl = "https://bittrex.com/api/v1.1/account/getbalances?apikey=API_KEY"

type balResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	Balances []struct {
		Currency  string  `json:"currency"`
		Total     float64 `json:"balance"`
		Available float64 `json:"available"`
		Locked    float64 `json:"pending"`
	} `json:"result"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, balancesUrl, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	var r balResp
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	balances = make(crypto.Balances)
	for _, b := range r.Balances {
		err := balances.Add(b.Currency, b.Available, b.Total, b.Locked)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", b.Currency, err)
		}
	}

	return
}
