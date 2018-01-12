package gdax

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "gdax.com"

var log = utils.Log().WithField("exchange", Domain)

// DOCS: https://docs.gdax.com/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "gdax",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
