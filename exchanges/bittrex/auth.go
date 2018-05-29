package bittrex

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"net/url"
		"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func privateRequest(c crypto.Credentials, url2 string, params map[string]string) (response *http.Response, err error) {
	nonce := fmt.Sprintf("%d", time.Now().Unix())

	v := url.Values{}
	v.Add("apikey", c.Key)
	v.Add("nonce", nonce)

	for key, val := range params {
		v.Add(key, val)
	}

	// Always a GET, according to docs…
	req, err := http.NewRequest("GET", url2, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = v.Encode()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("apisign", utils.HmacSign(sha512.New, req.URL.String(), c.Secret))

	return utils.NetClient().Do(req)
}
