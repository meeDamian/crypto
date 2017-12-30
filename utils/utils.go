package utils

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func NetClient() *http.Client {
	return &http.Client{Timeout: 15 * time.Second}
}

func ToFloat(v interface{}) (float64, error) {
	switch t := v.(type) {
	case string:
		return strconv.ParseFloat(t, 64)

	case float64:
		return t, nil
	}

	return 0, errors.Errorf("can't convert %v to float64: unsupported type", v)
}
