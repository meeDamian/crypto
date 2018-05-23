package yobit

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const apiUrl = "https://yobit.net/tapi/"

func privateRequest(c crypto.Credentials, apiMethod string, params map[string]string) (response *http.Response, err error) {
	v := url.Values{}
	v.Add("method", apiMethod)
	v.Add("nonce", fmt.Sprintf("%d", time.Now().Unix()))

	for key, val := range params {
		v.Add(key, val)
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Key", c.Key)
	req.Header.Add("Sign", utils.HmacSign(sha512.New, v.Encode(), c.Secret))

	return utils.NetClient().Do(req)
}

