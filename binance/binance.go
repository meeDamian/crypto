package binance

import "github.com/meeDamian/crypto"

const Domain = "binance.com"

// DOCS: https://www.binance.com/restapipub.html
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "Binance",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
