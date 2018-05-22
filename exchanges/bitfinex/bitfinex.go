package bitfinex

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "bitfinex.com"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Iot}
)

// DOCS: https://docs.bitfinex.com/v1/reference
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Bitfinex",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
