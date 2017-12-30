package acx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func signature(method, url, query, secret string) string {
	canonical := strings.Split(url, Domain)[1]

	toSign := fmt.Sprintf("%s|%s|%s",
		strings.ToUpper(method),
		strings.ToLower(canonical),
		query,
	)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(toSign))

	return hex.EncodeToString(mac.Sum(nil))
}

func privateRequest(c crypto.Credentials, method, url string, params map[string]string) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("access_key", c.Key)
	q.Add("tonce", fmt.Sprintf("%d", time.Now().Unix()*1e3))

	for key, val := range params {
		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()

	signature := signature(method, url, req.URL.RawQuery, c.Secret)

	req.URL.RawQuery += fmt.Sprintf("&signature=%s", signature)

	return utils.NetClient().Do(req)
}
