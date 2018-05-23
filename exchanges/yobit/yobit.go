package yobit

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "yobit.net"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Bcc}
)

// DOCS: https://yobit.io/en/api/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Yobit",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
