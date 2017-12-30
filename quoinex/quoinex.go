package quoinex

import (
	"github.com/meeDamian/crypto"
)

const Domain = "quoinex.com"

// DOCS: https://developers.quoine.com/v2
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "QUOINEX",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
