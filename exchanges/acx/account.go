package acx

import (
	"encoding/json"
	"strconv"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
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
		return crypto.Balances{}, err
	}

	defer res.Body.Close()

	var m me
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		err = errors.Wrapf(err, "can't decode me from %s", Domain)
		return
	}

	balances = make(crypto.Balances)
	for _, b := range m.Accounts {
		currency, err := currencies.Get(b.Currency)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: unknown currency", b.Currency)
			continue
		}

		balance, err := strconv.ParseFloat(b.Balance, 64)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: can't convert Balance=%s to float", b.Currency, b.Balance)
			continue
		}

		locked, err := strconv.ParseFloat(b.Locked, 64)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: can't convert Locked=%s to float", b.Currency, b.Locked)
			continue
		}

		balances[currency.Name] = crypto.Balance{
			Available: balance - locked,
			Total:     balance,
		}
	}

	return

}
