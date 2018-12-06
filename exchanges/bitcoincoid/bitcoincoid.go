package bitcoincoid

import (
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/markets"
	"github.com/meeDamian/crypto/utils"
)

const Domain = "bitcoin.co.id"

var (
	log = utils.Log().WithField("exchange", Domain)

	aliases = []string{
		currencies.Nem, // Xem
		currencies.Drk, // Dash
		currencies.Str, // Xlm
	}
)

// DOCS: https://vip.bitcoin.co.id/downloads/BITCOINCOID-API-DOCUMENTATION.pdf
func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Bitcoin Indonesia",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   func() ([]markets.Market, error) { return marketList, nil },

		// private
		Balances: Balances,
	}
}
