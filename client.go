package crypto

import (
	"github.com/meeDamian/crypto/orderbook"
	"github.com/pkg/errors"
)

const (
	methodNotImplemented = "%s.%s() not implemented"
	credentialsMissing   = "credentials for %s.%s() not provided"
)

func handleNilMarkets(e Exchange) func() ([]Market, error) {
	if e.Markets != nil {
		return e.Markets
	}

	return func() ([]Market, error) {
		return []Market{}, errors.Errorf(methodNotImplemented, e.Domain, "Markets")
	}
}

func handleNilOrderBook(e Exchange) func(Market) (orderbook.OrderBook, error) {
	if e.OrderBook != nil {
		return e.OrderBook
	}

	return func(market Market) (orderbook.OrderBook, error) {
		return orderbook.OrderBook{}, errors.Errorf(methodNotImplemented, e.Domain, "OrderBook")
	}
}

func handleNilAllOrderBooks(e Exchange) func() (map[Market]orderbook.OrderBook, error) {
	if e.AllOrderBooks != nil {
		return e.AllOrderBooks
	}

	return func() (map[Market]orderbook.OrderBook, error) {
		return map[Market]orderbook.OrderBook{}, errors.Errorf(methodNotImplemented, e.Domain, "AllOrderBooks")
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
