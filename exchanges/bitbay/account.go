package bitbay

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

type balResp struct {
	Balances map[string]struct {
		Available string `json:"available"`
		Locked    string `json:"locked"`
	} `json:"balances"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "info", nil)
	if err != nil {
		return balances, err
	}

	defer res.Body.Close()

	var bal balResp
	err = json.NewDecoder(res.Body).Decode(&bal)
	if err != nil {
		return balances, errors.Wrap(err, "can't json-decode response")
	}

	balances = make(crypto.Balances)
	for name, b := range bal.Balances {
		err := balances.Add(name, b.Available, nil, b.Locked)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", name, err)
		}
	}

	return
}
