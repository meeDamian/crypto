package bitbay

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "bitbay.net"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Bcc}
)

// DOCS: https://bitbay.net/en/public-api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "BitBay",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   func() ([]crypto.Market, error) { return marketList, nil },

		// private
		Balances: Balances,
	}
}
