package yobit

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

type balResp struct {
	Success int     `json:"success"`
	Error   *string `json:"error"`
	Return *struct {
		Available map[string]float64 `json:"funds"`
		Total     map[string]float64 `json:"funds_incl_orders"`
	} `json:"return"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "getInfo", nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var info balResp
	err = json.NewDecoder(res.Body).Decode(&info)
	if err != nil {
		return
	}

	if info.Success == 0 {
		return nil, errors.Errorf("error downloading balances: %s", *info.Error)
	}

	balances = make(crypto.Balances)
	for name, total := range info.Return.Total {
		var available *float64
		if val, ok := info.Return.Available[name]; ok {
			available = &val
		}

		err := balances.Add(name, *available, total, nil)
		if err != nil {
			log.Debugf("skipping balance of %s = %f/%f: %v", name, *available, total, err)
		}
	}

	return
}
