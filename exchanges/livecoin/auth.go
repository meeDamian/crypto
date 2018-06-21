package livecoin

import (
	"crypto/sha256"
	"net/http"
	"net/url"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func privateRequest(c crypto.Credentials, method, url2 string, params map[string]string) (response *http.Response, err error) {
	v := url.Values{}
	for key, val := range params {
		v.Add(key, val)
	}

	req, err := http.NewRequest(method, url2, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("API-Key", c.Key)
	req.Header.Set("Sign", strings.ToUpper(utils.HmacSign(sha256.New, v.Encode(), c.Secret)))

	return utils.NetClient().Do(req)
}
