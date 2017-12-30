package bitcoincoid

import "github.com/meeDamian/crypto"

const Domain = "bitcoin.co.id"

// DOCS: https://vip.bitcoin.co.id/downloads/BITCOINCOID-API-DOCUMENTATION.pdf
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Bx",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   func() ([]crypto.Market, error) { return marketList, nil },

		// private
		Balances: Balances,
	}
}
