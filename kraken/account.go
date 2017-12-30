package kraken

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

const balUrl = "https://api.kraken.com/0/private/Balance"

type balResp struct {
	Result map[string]string `json:"result"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, balUrl, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var bal balResp
	err = json.NewDecoder(res.Body).Decode(&bal)
	if err != nil {
		err = errors.Wrapf(err, "can't decode me from %s", Domain)
		return
	}

	balances = make(crypto.Balances)
	for name, balance := range bal.Result {
		code, err := removeKrakenNonsense(name)
		if err != nil {
			log.Println("EEE", err)
			continue
		}

		bal, err := strconv.ParseFloat(balance, 64)
		if err != nil {
			crypto.Log().Debugf("skipping balance of %s: can't convert Balance=%s to float", name, balance)
			continue
		}

		balances[code] = crypto.Balance{
			Available: bal,
			Total:     bal,
		}
	}

	return
}
