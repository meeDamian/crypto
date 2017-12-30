package coinfloor

import (
	"github.com/meeDamian/crypto"
)

const Domain = "coinfloor.co.uk"

// DOCS: https://github.com/coinfloor/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "coinfloor",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   func() ([]crypto.Market, error) { return marketList, nil },
	}
}
