package acx

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "acx.io"

var log = utils.Log().WithField("exchange", Domain)

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
