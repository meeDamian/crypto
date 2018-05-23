package bitfinex

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

type (
	balance struct {
		Type,
		Currency,
		Amount,
		Available string
	}

	aggregatedBalance struct {
		Total, Available float64
	}
)

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "POST", "balances")
	if err != nil {
		return balances, err
	}

	defer res.Body.Close()

	var bals []balance
	err = json.NewDecoder(res.Body).Decode(&bals)
	if err != nil {
		return balances, errors.Wrap(err, "can't json-decode response")
	}

	tmpBalances := make(map[string]aggregatedBalance)
	for _, b := range bals {
		total, err := utils.ToFloat(b.Amount)
		if err != nil {
			log.Debugf("skipping total balance of %s (%s) = %s: %v", b.Currency, b.Type, b.Amount, err)
			continue
		}

		available, err := utils.ToFloat(b.Available)
		if err != nil {
			log.Debugf("skipping available balance of %s (%s) = %s: %v", b.Currency, b.Type, b.Available, err)
			continue
		}

		x := aggregatedBalance{
			Total:     total,
			Available: available,
		}

		if val, ok := tmpBalances[b.Currency]; ok {
			x.Available += val.Available
			x.Total += val.Total
		}

		tmpBalances[b.Currency] = x
	}

	balances = make(crypto.Balances)
	for name, b := range tmpBalances {
		err := balances.Add(name, b.Available, b.Total, nil)
		if err != nil {
			log.Debugf("skipping balance of %s: %v", name, err)
		}
	}

	return
}
