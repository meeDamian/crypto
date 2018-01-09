package kraken

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

const balancesUrl = "https://api.kraken.com/0/private/Balance"

type balResp struct {
	Result map[string]string `json:"result"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, balancesUrl, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var bal balResp
	err = json.NewDecoder(res.Body).Decode(&bal)
	if err != nil {
		return balances, errors.Wrap(err, "can't json-decode response")
	}

	balances = make(crypto.Balances)
	for currency, available := range bal.Result {
		currencyName, err := removeKrakenNonsense(currency)
		if err != nil {
			log.Debug(err)
			continue
		}

		err = balances.Add(currencyName, available, nil, nil)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", currencyName, err)
		}
	}

	return
}
