package nzbcx

import (
	"fmt"
	"strings"

	"github.com/meeDamian/crypto"
	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/pkg/errors"
)

const orderBookUrl = "https://nzbcx.com/api/orderbook/%s%s"

var marketList = []crypto.Market{
	{Btc, Nzd},
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, strings.ToLower(m.Asset), strings.ToLower(m.PricedIn))

	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrap(err, "unable to fetch Order Book")
	}

	return
}
