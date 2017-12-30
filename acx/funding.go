package acx

import (
	"encoding/json"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

const (
	withdrawalsUrl = "https://acx.io/api/v2/withdraws.json"
	depositsUrl    = "https://acx.io/api/v2/deposits.json"
)

type bla struct {
	Id         int64     `json:"id"`
	Currency   string    `json:"currency"`
	Amount     string    `json:"amount"`
	Fee        string    `json:"fee"`
	PlacedAt   time.Time `json:"created_at"`
	ExecutedAt time.Time `json:"done_at"`
	State      string    `json:"state"`

	TxId          *string `json:"txid"`
	Confirmations *string `json:"confirmations"`
	TxHash        *string `json:"transaction_hash"`
}

func Withdrawals(c crypto.Credentials) (blas []bla, err error) {
	res, err := privateRequest(c, "GET", withdrawalsUrl, map[string]string{
		"limit": "100",
	})

	if err != nil {
		return []bla{}, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&blas)
	if err != nil {
		err = errors.Wrapf(err, "can't decode withdrawals from %s", Domain)
	}

	return
}

func Deposits(c crypto.Credentials) (blas []bla, err error) {
	res, err := privateRequest(c, "GET", depositsUrl, map[string]string{
		"limit": "100",
	})

	if err != nil {
		return []bla{}, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&blas)
	if err != nil {
		err = errors.Wrapf(err, "can't decode deposits from %s", Domain)
	}

	return
}
