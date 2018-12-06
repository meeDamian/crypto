package crypto

import (
	"github.com/meeDamian/crypto/markets"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/pkg/errors"
)

const (
	methodNotImplemented = "%s.%s() not implemented"
	credentialsMissing   = "credentials for %s.%s() not provided"
)

func handleNilMarkets(e Exchange) func() ([]markets.Market, error) {
	if e.Markets != nil {
		return e.Markets
	}

	return func() ([]markets.Market, error) {
		return []markets.Market{}, errors.Errorf(methodNotImplemented, e.Domain, "Markets")
	}
}

func handleNilOrderBook(e Exchange) func(markets.Market) (orderbook.OrderBook, error) {
	if e.OrderBook != nil {
		return e.OrderBook
	}

	return func(market markets.Market) (orderbook.OrderBook, error) {
		return orderbook.OrderBook{}, errors.Errorf(methodNotImplemented, e.Domain, "OrderBook")
	}
}

func handleNilAllOrderBooks(e Exchange) func() (map[markets.Market]orderbook.OrderBook, error) {
	if e.AllOrderBooks != nil {
		return e.AllOrderBooks
	}

	return func() (map[markets.Market]orderbook.OrderBook, error) {
		return map[markets.Market]orderbook.OrderBook{}, errors.Errorf(methodNotImplemented, e.Domain, "AllOrderBooks")
	}
}

func processBalancesFn(e Exchange, c *Credentials) func() (Balances, error) {
	if e.Balances == nil {
		return func() (Balances, error) {
			return Balances{}, errors.Errorf(methodNotImplemented, e.Domain, "Balances")
		}
	}

	if c == nil {
		return func() (Balances, error) {
			return Balances{}, errors.Errorf(credentialsMissing, e.Domain, "Balances")
		}
	}

	return func() (Balances, error) {
		return e.Balances(*c)
	}
}

func Client(exchangeFn func() Exchange, c *Credentials) ExchangeClient {
	exchange := exchangeFn()

	return ExchangeClient{
		Domain: exchange.Domain,
		Name:   exchange.Name,

		Markets:       handleNilMarkets(exchange),
		OrderBook:     handleNilOrderBook(exchange),
		AllOrderBooks: handleNilAllOrderBooks(exchange),
		Balances:      processBalancesFn(exchange, c),
	}
}
