package kraken

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
)

const balancesUrl = "https://api.kraken.com/0/private/Balance"

type balResp struct {
	Result map[string]string `json:"result"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, balancesUrl, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	var bal balResp
	err = json.NewDecoder(res.Body).Decode(&bal)
	if err != nil {
		return
	}

	balances = make(crypto.Balances)
	for currency, available := range bal.Result {
		currencyName, err := removePrefix(currency)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", currency, err)
			continue
		}

		err = balances.Add(currencyName, available, nil, nil)
		if err != nil {
			log.Debugf("skipping balance of %s = %s: %v", currencyName, available, err)
		}
	}

	return
}
