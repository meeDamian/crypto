package acx

import (
	"github.com/meeDamian/crypto"
)

const Domain = "acx.io"

// DOCS: https://acx.io/documents/api_v2
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "acx",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
