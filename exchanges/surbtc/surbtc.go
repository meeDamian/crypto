package surbtc

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "surbtc.com"

var log = utils.Log().WithField("exchange", Domain)

// DOCS: http://api.surbtc.com/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "SURBTC",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   Markets,
	}
}
