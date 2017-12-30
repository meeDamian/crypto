package bitbay

import "github.com/meeDamian/crypto"

const Domain = "bitbay.net"

// DOCS: https://bitbay.net/en/public-api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "BitBay",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   func() ([]crypto.Market, error) { return marketList, nil },

		// private
		Balances: Balances,
	}
}
