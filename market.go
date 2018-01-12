package crypto

import (
	"fmt"

	"github.com/meeDamian/crypto/currencies"
	"github.com/pkg/errors"
)

type Market struct {
	Asset, PricedIn string
}

func (m Market) String() string {
	return fmt.Sprintf("%s/%s", m.Asset, m.PricedIn)
}

func NewMarket(asset, price string) (m Market, err error) {
	a, err := currencies.Get(asset)
	if err != nil {
		return m, err
	}

	p, err := currencies.Get(price)
	if err != nil {
		return m, err
	}

	return Market{a.Name, p.Name}, nil
}

// appends market only if both currencies are known
func AppendMarket(markets []Market, rawAsset, rawPrice string) ([]Market, error) {
	asset, err := currencies.Get(rawAsset)
	if err != nil {
		return markets, err
	}

	price, err := currencies.Get(rawPrice)
	if err != nil {
		return markets, err
	}

	return append(markets, Market{
		asset.Name,
		price.Name,
	}), nil
}

// extracts a valid market from a given symbol, or returns an error
func NewMarketFromSymbol(symbol string) (market Market, err error) {
	matches := pairRegExp.FindAllStringSubmatch(symbol, -1)
	if len(matches) == 0 {
		return market, errors.Errorf("symbol %s is invalid or contains unknown currency", symbol)
	}

	return NewMarket(matches[0][1], matches[0][2])
}
