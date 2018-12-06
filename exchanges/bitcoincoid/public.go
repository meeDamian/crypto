package bitcoincoid

import (
	"encoding/json"
	"fmt"
	"github.com/meeDamian/crypto/markets"
	"strings"

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

var marketList = []markets.Market{
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

func morph(name string) string {
	return strings.ToLower(currencies.Morph(name, aliases))
}

func OrderBook(m markets.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl, morph(m.Asset), morph(m.PricedIn))

	res, err := utils.NetClient().Get(url)
	if err != nil {
		return ob, err
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	return orderbook.Normalise(r.Asks, r.Bids)
}
