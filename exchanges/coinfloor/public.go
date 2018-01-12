package coinfloor

import (
	"fmt"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/meeDamian/crypto/orderbook"
)

const orderBookUrl = "https://webapi.coinfloor.co.uk:8090/bist/%s/%s/order_book/"

var (
	marketList = []crypto.Market{
		{Btc, Gbp},
		{Btc, Eur},
		{Btc, Pln},
		{Btc, Usd},

		{Bch, Gbp},

		// "coming soon"
		//{Eth, Gbp},
		//{Etc, Gbp},
		//{Xrp, Gbp},
		//{Ltc, Gbp},
	}
)

func morph(name string) string {
	return currencies.Morph(name, aliases)
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), m.PricedIn)
	return orderbook.Download(url)
}
