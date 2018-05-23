package crypto

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
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
	pairRegExp = *regexp.MustCompile(fmt.Sprintf(`(?i)^[ZX]?(%[1]s)\/?[_ZX]?(%[1]s)$`, orSymbols))
}
