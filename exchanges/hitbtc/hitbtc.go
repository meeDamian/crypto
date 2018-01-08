package hitbtc

import "github.com/meeDamian/crypto"

const Domain = "hitbtc.com"

// DOCS: https://api.hitbtc.com/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "HitBTC",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}

