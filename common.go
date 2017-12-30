package crypto

import "github.com/meeDamian/crypto/orderbook"

type (
	Exchange struct {
		Name, Domain string

		// public
		OrderBook func(Market) (orderbook.OrderBook, error)
		Markets   func() ([]Market, error)

		// private
		Balances func(Credentials) (Balances, error)
		//Orders func() ([]orderbook.Order, error)
	}

	Credentials struct {
		Id *string // because Bitstamp…

		Name, Key, Secret string
	}
)
