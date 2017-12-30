package crypto

import (
	"github.com/meeDamian/crypto/currencies"
	"github.com/pkg/errors"
	"github.com/meeDamian/crypto/utils"
)

type (
	Balance struct {
		Available,
		Total,
		Locked float64
	}

	Balances map[string]Balance
)

// NOTE: available = total - locked
func (b *Balances) Add(currencyName string, available, total, locked interface{}) error {
	currency, err := currencies.Get(currencyName)
	if err != nil {
		return errors.Errorf("unknown currency: %s", currencyName)
	}

	if available == nil {
		return errors.Errorf("amount of 'available' %s not provided", currencyName)
	}

	var balance Balance
	balance.Available, err = utils.ToFloat(available)
	if err != nil {
		return errors.Wrap(err, "unable to parse 'available' amount")
	}

	if total != nil {
		balance.Total, err = utils.ToFloat(total)
		if err != nil {
			return errors.Wrap(err, "unable to parse 'total' amount")
		}

		if locked == nil {
			balance.Locked = balance.Total - balance.Available
		}
	}

	if locked != nil {
		balance.Locked, err = utils.ToFloat(locked)
		if err != nil {
			return errors.Wrap(err, "unable to parse 'locked' amount")
		}

		if total == nil {
			balance.Total = balance.Available + balance.Locked
		}
	}

	if total == nil && locked == nil {
		balance.Total = balance.Available
		balance.Locked = 0
	}

	(*b)[currency.Name] = balance

	return nil
}
