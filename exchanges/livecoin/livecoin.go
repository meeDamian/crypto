package livecoin

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "livecoin.net"

var log = utils.Log().WithField("exchange", Domain)

// DOCS: https://liqui.io/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Livecoin",
		Domain: Domain,

		// public
		Markets:       Markets,
		OrderBook:     OrderBook,
		AllOrderBooks: AllOrderBooks,

		// private
		//Balances: Balances,
	}
}
