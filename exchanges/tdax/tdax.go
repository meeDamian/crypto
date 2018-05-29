package tdax

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "tdax.com"

var log = utils.Log().WithField("exchange", Domain)

// DOCS: https://api-docs.tdax.com/apis/public/orders.html
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "TDAX",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		//private
		Balances: Balances,
	}
}
