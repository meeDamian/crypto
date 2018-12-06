package coinfloor

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/markets"
)

const Domain = "coinfloor.co.uk"

var aliases = []string{currencies.Xbt}

// DOCS: https://github.com/coinfloor/api
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "CoinFloor",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   func() ([]markets.Market, error) { return marketList, nil },
	}
}
