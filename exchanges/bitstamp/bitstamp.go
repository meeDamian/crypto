package bitstamp

import (
	"github.com/meeDamian/crypto"
)

const Domain = "bitstamp.net"

// DOCS: https://www.bitstamp.net/api/v2/trading-pairs-info/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Bitstamp",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
