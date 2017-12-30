package luno

import (
	"github.com/meeDamian/crypto"
)

const Domain = "luno.com"

// DOCS: https://www.luno.com/en/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "Luno",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
