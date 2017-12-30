package gdax

import "github.com/meeDamian/crypto"

const Domain = "gdax.com"

// DOCS: https://docs.gdax.com/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "gdax",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
