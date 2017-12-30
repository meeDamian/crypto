package bitstamp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func signature(c crypto.Credentials, nonce string) string {
	mac := hmac.New(sha256.New, []byte(c.Secret))
	mac.Write([]byte(nonce + *c.Id + c.Key))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
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
