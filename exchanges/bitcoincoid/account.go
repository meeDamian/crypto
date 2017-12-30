package bitcoincoid

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/pkg/errors"
)

type balResp struct {
	Return struct {
		Balances map[string]interface{} `json:"balance"`
	} `json:"return"`
}

func magic(n interface{}) (float64, error) {
	switch v := n.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case float64:
		return v, nil
	case int64:
		return float64(v), nil
	}

	return 0, errors.Errorf("can't convert %s to float", n)
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "getInfo", nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	var r balResp
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		err = errors.Wrapf(err, "can't decode me from %s", Domain)
		return
	}

	balances = make(crypto.Balances)
	for name, b := range r.Return.Balances {
		currency, err := currencies.Get(name)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: unknown currency", name)
			continue
		}

		balance, err := magic(b)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: %v", name, err)
			continue
		}

		balances[currency.Name] = crypto.Balance{
			Available: balance,
			Total:     balance,
		}
	}

	return
}
