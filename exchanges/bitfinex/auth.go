package bitfinex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const apiUrl = "https://api.bitfinex.com/v1/%s"

func signature(c crypto.Credentials, toSign string) string {
	mac := hmac.New(sha512.New384, []byte(c.Secret))
	mac.Write([]byte(toSign))

	return hex.EncodeToString(mac.Sum(nil))
}

func privateRequest(c crypto.Credentials, method, apiMethod string) (response *http.Response, err error) {
	nonce := fmt.Sprintf("%d", time.Now().Unix())

	payload := map[string]interface{}{
		"request": "/v1/" + apiMethod,
		"nonce":   nonce,
	}

	req, err := http.NewRequest(method, fmt.Sprintf(apiUrl, apiMethod), nil)
	if err != nil {
		return nil, err
	}

	p, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(p)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-BFX-APIKEY", c.Key)
	req.Header.Add("X-BFX-PAYLOAD", encoded)
	req.Header.Add("X-BFX-SIGNATURE", signature(c, encoded))

	return utils.NetClient().Do(req)
}
