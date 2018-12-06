package crypto

import (
	"github.com/meeDamian/crypto/markets"
	"github.com/meeDamian/crypto/orderbook"
)

type (
	Credentials struct {
		// human-readable Name of the account
		Name string

		// API Key & Secret
		Key, Secret string

		// Id needed by bitstamp and tdax
		Id *string
	}

	Exchange struct {
		// each exchange SHOULD specify its human-readable name
		Name string

		// each exchange MUST specify its generic top level domain (`www.`, etc should be omitted)
		Domain string

		/**
		 * public
		**/
		// returns a list of all Markets on available on a given exchange. Includes disabled Markets.
		//      Limited to supported currencies only, see currencies/currencies.go and currencies/symbols/symbols.go for more
		Markets func() ([]markets.Market, error)

		// returns OrderBook for requested Market or error
		OrderBook func(markets.Market) (orderbook.OrderBook, error)

		/**
		 * private
		**/
		// returns all exchange Balances (Total, Available, Locked) for account credentials provided
		Balances func(Credentials) (Balances, error)

		/**
		 * optional
		**/
		// returns OrderBooks of ALL available markets. Should only be implemented if a "shortcut" endpoint exists
		//      If only some markets couldn't be downloaded, error should be logged, but not returned.
		//      Error only if no usable data can be returned
		AllOrderBooks func() (map[markets.Market]orderbook.OrderBook, error)
	}

	// The same as Exchange, except:
	//      1) Doesn't require explicit Credentials passed each time
	//      2) Inserts methods returning "not implemented" error in place of missing ones
	//      3) (TODO) all methods respect rate-limiting; scoped to client instance(!)
	ExchangeClient struct {
		Name          string
		Domain        string
		Markets       func() ([]markets.Market, error)
		OrderBook     func(markets.Market) (orderbook.OrderBook, error)
		Balances      func() (Balances, error)
		AllOrderBooks func() (map[markets.Market]orderbook.OrderBook, error)
	}
)
