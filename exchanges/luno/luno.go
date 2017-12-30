package luno

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
	"github.com/meeDamian/crypto/currencies"
)

const Domain = "luno.com"

var (
	log     = utils.Log().WithField("exchange", Domain)
	aliases = []string{currencies.Xbt}
)

// DOCS: https://www.luno.com/en/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "Luno",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
