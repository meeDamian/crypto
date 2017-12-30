package bitstamp

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/pkg/errors"
)

const balancesUrl = "https://www.bitstamp.net/api/v2/balance/"

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "POST", balancesUrl, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	var r map[string]string
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		err = errors.Wrapf(err, "can't decode me from %s", Domain)
		return
	}

	balances = make(crypto.Balances)
	for name, amount := range r {
		b := strings.Split(name, "_")
		if len(b) < 2 || b[1] != "available" {
			continue
		}

		currency, err := currencies.Get(b[0])
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: unknown currency", name)
			continue
		}

		a, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			crypto.Log().Debugf("skipping 'available' balance of %s: %v", name, err)
			continue
		}

		balances[currency.Name] = crypto.Balance{
			Available: a,
			Total:     a,
		}
	}

	return
}
