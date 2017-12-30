package crypto

import (
	"github.com/meeDamian/crypto/orderbook"
	"strings"
	"regexp"
	"fmt"
	"github.com/meeDamian/crypto/currencies"
)

type (
	Exchange struct {
		Name, Domain string

		// public
		OrderBook func(Market) (orderbook.OrderBook, error)
		Markets   func() ([]Market, error)

		// private
		Balances func(Credentials) (Balances, error)
	}

	Credentials struct {
		Id *string // because Bitstamp…

		Name, Key, Secret string
	}
)

var pairRegExp regexp.Regexp

func init() {
	var symbols []string
	for symbol := range currencies.All() {
		symbols = append(symbols, symbol)
	}

	orSymbols := strings.Join(symbols, "|")
	pairRegExp = *regexp.MustCompile(fmt.Sprintf(`^[ZX]?(%[1]s)\/?[ZX]?(%[1]s)$`, orSymbols))
}
