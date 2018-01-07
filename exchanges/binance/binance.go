package binance

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
	"github.com/meeDamian/crypto/currencies"
)

const Domain = "binance.com"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Bcc}
)

// DOCS: https://www.binance.com/restapipub.html
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "Binance",
		Domain:    Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
