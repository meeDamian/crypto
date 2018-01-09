package hitbtc

import (
	"net/http"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func privateRequest(c crypto.Credentials, method, url string, params map[string]string) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	for key, val := range params {
		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(c.Key, c.Secret)

	return utils.NetClient().Do(req)
}
