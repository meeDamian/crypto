package kraken

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "kraken.com"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Xbt, currencies.Xdg}
)

// DOCS: https://www.kraken.com/help/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Kraken",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
