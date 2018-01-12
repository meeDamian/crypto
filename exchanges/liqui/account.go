package liqui

import (
	"encoding/json"

	"github.com/meeDamian/crypto"
	"github.com/pkg/errors"
)

type infoResp struct {
	Success int     `json:"success"`
	Error   *string `json:"error"`
	Return *struct {
		Funds map[string]float64 `json:"funds"`
	} `json:"return"`
}

func Balances(c crypto.Credentials) (balances crypto.Balances, err error) {
	res, err := privateRequest(c, "getInfo", nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var info infoResp
	err = json.NewDecoder(res.Body).Decode(&info)
	if err != nil {
		return
	}

	if info.Success == 0 {
		return nil, errors.Errorf("error downloading balances: %s", *info.Error)
	}

	balances = make(crypto.Balances)
	for name, b := range info.Return.Funds {
		err := balances.Add(name, b, nil, nil)
		if err != nil {
			log.Debugf("skipping balance of %s = %f: %v", name, b, err)
		}
	}

	return
}
