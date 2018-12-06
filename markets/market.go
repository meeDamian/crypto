package markets

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/meeDamian/crypto/currencies"
	"github.com/pkg/errors"
)

type Market struct {
	Asset, PricedIn string
}

var pairRegExp regexp.Regexp

func init() {
	var symbols []string
	for symbol := range currencies.All() {
		symbols = append(symbols, symbol)
	}

	orSymbols := strings.Join(symbols, "|")
	pairRegExp = *regexp.MustCompile(fmt.Sprintf(`(?i)^[ZX]?(%[1]s)\/?[_ZX]?(%[1]s)$`, orSymbols))
}

func (m Market) String() string {
	return fmt.Sprintf("%s/%s", m.Asset, m.PricedIn)
}

var currencyNotSupportedCustomTrigger = func(string, ...error) {} // noop by default

func SetCurrencyNotSupportedTrigger(userFn func(string, ...error)) {
	if userFn == nil {
		return
	}

	currencyNotSupportedCustomTrigger = userFn
}

func splitSymbol(symbol string, fn func(string, ...error)) {
	codes := strings.Split(symbol, "_")
	if len(codes) > 1 {
		New(codes[0], codes[1])
		return
	}

	codes = strings.Split(symbol, "/")
	if len(codes) > 1 {
		New(codes[0], codes[1])
		return
	}

	fn("", errors.Errorf("Unable to split %s into currencies", symbol))
}

func New(asset, price string) (m Market, error error) {
	a, err := currencies.Get(asset)
	if err != nil {
		currencyNotSupportedCustomTrigger(asset)
		error = err
	}

	p, err := currencies.Get(price)
	if err != nil {
		currencyNotSupportedCustomTrigger(price)
		error = err
	}

	if error != nil {
		return m, error
	}

	return Market{a.Name, p.Name}, nil
}

// appends market only if both currencies are known
func Append(markets []Market, rawAsset, rawPrice string) ([]Market, error) {
	market, err := New(rawAsset, rawPrice)
	if err != nil {
		return markets, err
	}

	return append(markets, market), nil
}

// extracts a valid market from a given symbol, or returns an error
func NewFromSymbol(symbol string) (market Market, err error) {
	matches := pairRegExp.FindAllStringSubmatch(symbol, -1)
	if len(matches) == 0 {
		splitSymbol(symbol, currencyNotSupportedCustomTrigger)
		return market, errors.Errorf("symbol %s is invalid or contains unknown currency", symbol)
	}

	return New(matches[0][1], matches[0][2])
}
