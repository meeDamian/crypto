package hitbtc

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "hitbtc.com"

var log = utils.Log().WithField("exchange", Domain)

// DOCS: https://api.hitbtc.com/
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "HitBTC",
		Domain:    Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}

