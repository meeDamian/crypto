package bitbay

import (
	"encoding/json"
	"strconv"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
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
		return
	}

	defer res.Body.Close()

	var bal balResp
	err = json.NewDecoder(res.Body).Decode(&bal)
	if err != nil {
		err = errors.Wrapf(err, "can't decode me from %s", Domain)
		return
	}

	balances = make(crypto.Balances)
	for name, b := range bal.Balances {
		currency, err := currencies.Get(name)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: unknown currency", name)
			continue
		}

		available, err := strconv.ParseFloat(b.Available, 64)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: can't convert Balance=%s to float", name, b.Available)
			continue
		}

		locked, err := strconv.ParseFloat(b.Locked, 64)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: can't convert Locked=%s to float", name, b.Locked)
			continue
		}

		balances[currency.Name] = crypto.Balance{
			Available: available,
			Total:     available + locked,
		}
	}

	return
}
