package bitbay

import (
	"io/ioutil"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func History(c crypto.Credentials) (me string, err error) {
	res, err := privateRequest(c, "history", map[string]string{"currency": "btc"})
	if err != nil {
		return
	}

	defer res.Body.Close()

	respBody, _ := ioutil.ReadAll(res.Body)

	utils.Log().Println(res.Status)
	utils.Log().Println(string(respBody))

	return
}
