package kraken

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

func signature(fullUrl, nonce, data, secret string) (string, error) {
	parsedUrl, err := url.Parse(fullUrl)
	if err != nil {
		return "", errors.Wrapf(err, "can't extract path from url = %s", fullUrl)
	}

	base64secret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", errors.Wrapf(err, "can't encode secret to base64")
	}

	noncedData := sha256.Sum256([]byte(nonce + data))

	fullData := append([]byte(parsedUrl.Path), noncedData[:]...)

	mac := hmac.New(sha512.New, base64secret)
	mac.Write(fullData)

	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func privateRequest(c crypto.Credentials, url2 string, params map[string]string) (response *http.Response, err error) {
	nonce := fmt.Sprintf("%d", time.Now().UnixNano())

	v := url.Values{}
	v.Add("nonce", nonce)

	for key, val := range params {
		v.Add(key, val)
	}

	postData := v.Encode()
	req, err := http.NewRequest("POST", url2, strings.NewReader(postData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("API-Key", c.Key)

	sig, err := signature(url2, nonce, postData, c.Secret)
	if err != nil {
		return nil, err
	}

	req.Header.Add("API-Sign", sig)

	return utils.NetClient().Do(req)
}
