package crypto

import (
	"fmt"
	"strings"

	"github.com/meeDamian/crypto/currencies"
)

type Market struct {
	Asset, PricedIn string
}

func (m Market) ToUpper() Market {
	m.Asset = strings.ToUpper(m.Asset)
	m.PricedIn = strings.ToUpper(m.PricedIn)
	return m
}

func (m Market) String() string {
	return fmt.Sprintf("%s/%s", m.Asset, m.PricedIn)
}

func normaliseCurrency(name string) string {
	currency, err := currencies.Get(strings.ToUpper(name))
	if err != nil {
		Log().Debugf("Unknown currency %s left unchanged", name)
		return name
	}

	return currency.Name
}

func NewMarket(asset, price string) Market {
	return Market{
		strings.ToUpper(normaliseCurrency(asset)),
		strings.ToUpper(normaliseCurrency(price)),
	}
}
