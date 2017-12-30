package bx

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func signature(c crypto.Credentials, nonce string) string {
	uniqueSecret := fmt.Sprintf("%s%s%s", c.Key, nonce, c.Secret)
	x := sha256.Sum256([]byte(uniqueSecret))
	return fmt.Sprintf("%x", x)
}

func privateRequest(c crypto.Credentials, method, url2 string, params map[string]string) (response *http.Response, err error) {
	nonce := fmt.Sprintf("%d", time.Now().Unix())

	v := url.Values{}
	v.Add("key", c.Key)
	v.Add("nonce", nonce)
	v.Add("signature", signature(c, nonce))

	for key, val := range params {
		v.Add(key, val)
	}

	req, err := http.NewRequest(method, url2, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return utils.NetClient().Do(req)
}
