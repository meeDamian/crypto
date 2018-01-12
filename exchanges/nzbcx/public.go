package nzbcx

import (
	"fmt"
	"strings"

	"github.com/meeDamian/crypto"
	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/meeDamian/crypto/orderbook"
)

const orderBookUrl = "https://nzbcx.com/api/orderbook/%s%s"

var marketList = []crypto.Market{
	{Btc, Nzd},
}

func morph(name string) string {
	return strings.ToLower(name)
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), morph(m.PricedIn))
	return orderbook.Download(url)
}
