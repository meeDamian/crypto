package coinfloor

import (
	"fmt"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/pkg/errors"
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

	aliases = []string{currencies.Xbt}
)

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl,
		currencies.Morph(m.Asset, aliases),
		m.PricedIn,
	)

	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrapf(err, "unable to fetch %s Order Book", Domain)
	}

	return
}
