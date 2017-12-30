package bittrex

import "github.com/meeDamian/crypto"

const Domain = "bittrex.com"

// DOCS: https://bittrex.com/home/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "bittrex",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}

