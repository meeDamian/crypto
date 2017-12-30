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

func NewMarket(asset, price string) Market {
	return Market{
		currencies.Normalise(asset),
		currencies.Normalise(price),
	}
}

func NewMarketFromSymbol(symbol string) (market Market, err error) {
	matches := pairRegExp.FindAllStringSubmatch(symbol, -1)
	if len(matches) == 0 {
		err = errors.Errorf("symbol %s is either invalid or at least one of the currencies is unknown", symbol)
		return
	}

	return NewMarket(matches[0][1], matches[0][2]), nil
}
