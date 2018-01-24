package utils

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

func HmacSign(fn func() hash.Hash, data, secret string) string {
	mac := hmac.New(fn, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
