package kraken

import (
	"github.com/meeDamian/crypto"
)

const Domain = "kraken.com"

// DOCS: https://www.kraken.com/help/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "Kraken",
		Domain:    Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
