package surbtc

import "github.com/meeDamian/crypto"

const Domain = "surbtc.com"

// DOCS: http://api.surbtc.com/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "SURBTC",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
