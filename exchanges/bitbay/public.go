package bitbay

import (
	"fmt"
	"github.com/meeDamian/crypto/markets"

	"github.com/meeDamian/crypto/currencies"
	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/meeDamian/crypto/orderbook"
)

const orderBookUrl = "https://bitbay.net/API/Public/%s%s/orderbook.json"

var (
	marketList = []markets.Market{
		{Btc, Usd}, {Btc, Eur}, {Btc, Pln},
		{Ltc, Usd}, {Ltc, Eur}, {Ltc, Pln}, {Ltc, Btc},
		{Bch, Usd}, {Bch, Eur}, {Bch, Pln}, {Bch, Btc},
		{Eth, Usd}, {Eth, Eur}, {Eth, Pln}, {Eth, Btc},
		{Lsk, Usd}, {Lsk, Eur}, {Lsk, Pln}, {Lsk, Btc},
	}
)

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl,
		currencies.Morph(m.Asset, aliases),
		currencies.Morph(m.PricedIn, aliases),
	)

	return orderbook.Download(url)
}
