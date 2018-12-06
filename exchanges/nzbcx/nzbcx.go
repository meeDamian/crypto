package nzbcx

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/markets"
)

const Domain = "nzbcx.com"

// DOCS: https://nzbcx.com/docs/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "NZBCX",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   func() ([]markets.Market, error) { return marketList, nil },
	}
}
