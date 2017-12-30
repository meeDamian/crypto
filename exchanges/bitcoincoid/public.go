package bitcoincoid

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	. "github.com/meeDamian/crypto/currencies/symbols"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
)

const orderBookUrl = "https://vip.bitcoin.co.id/api/%s_%s/depth"

type obResponse struct {
	Bids []interface{} `json:"buy"`
	Asks []interface{} `json:"sell"`
}

var (
	marketList = []crypto.Market{
		{Bts, Btc},
		{Dash, Btc},
		{Doge, Btc},
		{Xem, Btc},

		{Bch, Idr},
		{Btc, Idr},
		{Btg, Idr},
		{Etc, Idr},
		{Waves, Idr},
		{Xzc, Idr},
		{Eth, Idr}, {Eth, Btc},
		{Ltc, Idr}, {Ltc, Btc},
		{Nxt, Idr}, {Nxt, Btc},
		{Xrp, Idr}, {Xrp, Btc},
		{Xlm, Idr}, {Xlm, Btc},
	}

	aliases = []string{
		currencies.Nem, // Xem
		currencies.Drk, // Dash
		currencies.Str, // Xlm
	}
)

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl,
		strings.ToLower(currencies.Morph(m.Asset, aliases)),
		strings.ToLower(currencies.Morph(m.PricedIn, aliases)),
	)

	fmt.Println(url)

	res, err := utils.NetClient().Get(url)
	if err != nil {
		return orderbook.OrderBook{}, err
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return orderbook.OrderBook{}, err
	}

	return orderbook.Normalise(r.Asks, r.Bids)
}
