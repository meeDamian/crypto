package main

import (
	"fmt"
	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/exchanges"
	"os"
	"strings"
)

var (
	exchange *crypto.Exchange
	symbols  []currencies.Currency
)

func verifyExchange(name string) *crypto.Exchange {
	for _, e := range exchanges.All {
		if e.Domain == name {
			return &e
		}
	}

	return nil
}

func init() {
	params := os.Args[1:]

	if len(params) == 0 {
		panic("at the very least exchange domain needs to be provided")
	}

	var exchangeName string
	if len(params) >= 1 {
		exchangeName = params[0]
	}

	exchange = verifyExchange(exchangeName)
	if exchange == nil {
		panic(fmt.Sprintf("invalid exchange name, valid ones are: \n\t%s", strings.Join(exchanges.AllDomains(), ", ")))
	}

	var symbolNames []string
	if len(params) > 1 {
		symbolNames = params[1:]
	}

	for _, s := range symbolNames {
		currency, err := currencies.Get(s)
		if err != nil {
			panic(fmt.Sprintf("invalid currency: %s", s))
		}

		symbols = append(symbols, currency)
	}
}

func main() {
	if len(symbols) == 0 {
		exchange.Markets()
	}

}
