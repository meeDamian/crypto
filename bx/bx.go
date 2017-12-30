package bx

import "github.com/meeDamian/crypto"

const Domain = "bx.in.th"

// DOCS: https://bx.in.th/info/api/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Bx",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
