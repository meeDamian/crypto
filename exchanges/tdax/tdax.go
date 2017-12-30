package tdax

import (
	"github.com/meeDamian/crypto"
)

const Domain = "tdax.com"

// DOCS: https://api-docs.tdax.com/apis/public/orders.html
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:      "TDAX",
		Domain:    Domain,
		OrderBook: OrderBook,
		Markets:   func() ([]crypto.Market, error) { return marketList, nil },
	}
}
