package bittrex

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "bittrex.com"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Bcc}
)

// DOCS: https://bittrex.com/home/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "bittrex",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
