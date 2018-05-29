package tdax

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"net/url"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func privateRequest(c crypto.Credentials, method, url2 string, params map[string]string) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url2, nil)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	//v.Add("Nonce", fmt.Sprintf("%d", time.Now().Unix()))
	for key, val := range params {
		v.Add(key, val)
	}

	req.URL.RawQuery = v.Encode()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("TDAX-API %s", c.Key))
	req.Header.Set("Signature", utils.HmacSign(sha512.New, v.Encode(), c.Secret))

	return utils.NetClient().Do(req)
}
