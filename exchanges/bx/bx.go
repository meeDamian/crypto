package bx

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "bx.in.th"

var log = utils.Log().WithField("exchange", Domain)

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
