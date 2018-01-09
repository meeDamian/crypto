package bitcoincoid

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

type balResp struct {
	Return struct {
		Balances       map[string]interface{} `json:"balance"`
		LockedBalanced map[string]interface{} `json:"balance_hold"`
	} `json:"return"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "getInfo", nil)
	if err != nil {
		return balances, err
	}

	defer res.Body.Close()

	var r balResp
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return balances, errors.Wrap(err, "can't json-decode response")
	}

	balances = make(crypto.Balances)
	for name, available := range r.Return.Balances {
		locked, ok := r.Return.LockedBalanced[name]
		if !ok {
			locked = 0
		}

		err := balances.Add(name, available, nil, locked)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", name, err)
		}
	}

	return
}
