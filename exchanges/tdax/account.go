package tdax

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

const meUrl = "https://api.tdax.com/users/%s"

type meResp struct {
	Balances map[string]float64 `json:"Balances"`
	Code     *string            `json:"code"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	url := fmt.Sprintf(meUrl, *c.Id)

	res, err := privateRequest(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var m meResp
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		if m.Code != nil {
			return balances, errors.New(*m.Code)
		}
	}

	err = currencyPrecisions()
	if err != nil {
		return crypto.Balances{}, errors.Wrapf(err, "can't download required currency precisions")
	}

	balances = make(crypto.Balances)
	for currency, rawBalance := range m.Balances {
		balance, err := normalize(currency, rawBalance)
		if err != nil {
			log.Debugf("skipping balance of %s = %f: %v", currency, rawBalance, err)
		}

		err = balances.Add(currency, balance, nil, nil)
		if err != nil {
			log.Debugf("skipping balance of %s = %f: %v", currency, balance, err)
		}
	}

	return
}
