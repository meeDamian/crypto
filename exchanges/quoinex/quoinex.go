package quoinex

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "quoinex.com"

var log = utils.Log().WithField("exchange", Domain)

// DOCS: https://developers.quoine.com/v2
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "QUOINEX",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
