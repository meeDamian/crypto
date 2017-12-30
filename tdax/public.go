package tdax

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const orderBookUrl = "https://api.tdax.com/orders?Symbol=%s_%s"

type (
	order struct {
		Price  float64 `json:"Price"`
		Volume float64 `json:"RemainQty"`
	}

	obResponse struct {
		Bids []order `json:"Bids"`
		Asks []order `json:"Asks"`
	}
)

var (
	marketList = []crypto.Market{
		{Btc, Thb},
		{Eth, Thb}, {Eth, Btc},
		{Btg, Thb}, {Btg, Btc}, {Btg, Eth},
		{Knc, Thb}, {Knc, Btc}, {Knc, Eth},
		{Ltc, Thb}, {Ltc, Btc}, {Ltc, Eth},
		{Neo, Thb}, {Neo, Btc}, {Neo, Eth},
		{Omg, Thb}, {Omg, Btc}, {Omg, Eth},
		{Xrp, Thb}, {Xrp, Btc}, {Xrp, Eth},
		{Xzc, Thb}, {Xzc, Btc}, {Xzc, Eth},
	}

	aliases = []string{currencies.Rpx}

	precisions = map[string]int{
		Thb: 2,
		Btc: 8,
		Eth: 18,
		Btg: 8,
		Ltc: 8,
		Xrp: 6,
	}
)

func normalisedPendingOrder(o order, m crypto.Market) orderbook.PendingOrder {
	volumePrecision, ok := precisions[m.Asset]
	if !ok {
		panic(errors.Errorf("precision of %s unknown", m.Asset))
	}

	pricePrecision, ok := precisions[m.PricedIn]
	if !ok {
		panic(errors.Errorf("precision of %s unknown", m.PricedIn))
	}

	return orderbook.PendingOrder{
		Price:  o.Price / math.Pow10(pricePrecision),
		Volume: o.Volume / math.Pow10(volumePrecision),
	}
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("unable to convert order book: %s", r)
		}
	}()

	url := fmt.Sprintf(orderBookUrl,
		currencies.Morph(m.Asset, aliases),
		currencies.Morph(m.PricedIn, aliases),
	)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		err = errors.Wrap(err, "unable to GET orderbook")
		return
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		err = errors.Wrap(err, "unable to decode response")
		return
	}

	var asks, bids []orderbook.PendingOrder
	for _, o := range r.Asks {
		asks = append(asks, normalisedPendingOrder(o, m))
	}

	for _, o := range r.Bids {
		bids = append(bids, normalisedPendingOrder(o, m))
	}

	ob, err = orderbook.Sort(orderbook.OrderBook{
		Asks: asks,
		Bids: bids,
	})

	if err != nil {
		err = errors.Wrapf(err, "unable to fetch %s Order Book", Domain)
	}

	return
}

func Markets() (_ []crypto.Market, err error) {
	return marketList, nil
}
