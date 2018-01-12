package bitstamp

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
)

const balancesUrl = "https://www.bitstamp.net/api/v2/balance/"

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "POST", balancesUrl, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var r map[string]string
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	balances = make(crypto.Balances)
	for name, amount := range r {
		b := strings.Split(name, "_")
		if len(b) < 2 || b[1] == "fee" {
			continue
		}

		currency, err := currencies.Get(b[0])
		if err != nil {
			log.Debugf("skipping balance of %s: unknown currency", name)
			continue
		}

		balance, ok := balances[currency.Name]
		if !ok {
			balance = crypto.Balance{}
		}

		a, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			log.Debugf("skipping '%s' balance of %s: %v", b[1], name, err)
			continue
		}

		switch b[1] {
		case "available":
			balance.Available = a
			break

		case "reserved":
			balance.Locked = a
			break

		case "balance":
			balance.Total = a
			break

		default:
			continue
		}

		balances[currency.Name] = balance
	}

	return
}
