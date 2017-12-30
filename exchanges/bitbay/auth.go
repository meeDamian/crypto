package bitbay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const apiUrl = "https://bitbay.net/API/Trading/tradingApi.php"

func signature(query, secret string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(query))

	return hex.EncodeToString(mac.Sum(nil))
}

func privateRequest(c crypto.Credentials, apiMethod string, params map[string]string) (response *http.Response, err error) {
	v := url.Values{}
	v.Add("method", apiMethod)
	v.Add("moment", fmt.Sprintf("%d", time.Now().Unix()))

	for key, val := range params {
		v.Add(key, val)
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("API-Key", c.Key)
	req.Header.Add("API-Hash", signature(v.Encode(), c.Secret))

	return utils.NetClient().Do(req)
}
