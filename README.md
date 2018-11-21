# `crypto` library

This library offers integrations with [some cryptocurrency exchanges]. 

[some cryptocurrency exchanges]: https://github.com/meeDamian/crypto/tree/master/exchanges  


# TODO: write more here.


### Defining new exchange interface



#### `exchange.go`

```go
package exchange

const Domain = "exchange.com"

func New() crypto.Exchange {
	return crypto.Exchange{
		Name:   "Exchange",
		Domain: Domain,

		// public
		OrderBook: OrderBook,
		Markets:   Markets,

		// private
		Balances: Balances,
	}
}
```

#### `public.go`

```Go
package exchange

func Markets() ([]crypto.Market, error) {}
func OrderBook(crypto.Market) (orderbook.OrderBook, error) {}
func AllOrderBooks() (map[string]orderbook.OrderBook, error) {}
```

#### `account.go`

```Go
package exchange

func Balances(crypto.Credentials) (crypto.Balances, error) {}
```
