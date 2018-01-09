package liqui

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "liqui.io"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Bcc}
)

// DOCS: https://liqui.io/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Liqui",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
