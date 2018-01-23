package tdax

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const (
	orderBookUrl  = "https://api.tdax.com/orders?Symbol=%s_%s"
	currenciesUrl = "https://api.tdax.com/public/getcurrencies"
	marketsUrl    = "https://api.tdax.com/public/getmarkets"
)

type (
	order struct {
		Price  float64 `json:"Price"`
		Volume float64 `json:"RemainQty"`
	}

	obResponse struct {
		Bids []order `json:"Bids"`
		Asks []order `json:"Asks"`
	}

	marketResponse []struct {
		Asset    string `json:"MarketCurrency"`
		PricedIn string `json:"BaseCurrency"`
	}

	currencyResponse []struct {
		Name    string `json:"Currency"`
		Divider int64  `json:"Divider"`
	}
)

var (
	marketList []crypto.Market
	precisions = make(map[string]int)

	aliases = []string{currencies.Rpx}
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

func currencyPrecisions() (err error) {
	res, err := utils.NetClient().Get(currenciesUrl)
	if err != nil {
		return
	}

	defer res.Body.Close()

	var cs currencyResponse
	err = json.NewDecoder(res.Body).Decode(&cs)
	if err != nil {
		return
	}

	for _, c := range cs {
		curr, err := currencies.Get(c.Name)
		if err != nil {
			log.Debugf("skipping precision of %s: %v", c.Name, err)
			continue
		}

		precisions[curr.Name] = int(math.Log10(float64(c.Divider)))
	}

	return
}

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	err = currencyPrecisions()
	if err != nil {
		return []crypto.Market{}, errors.Wrapf(err, "can't download required currency precisions")
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []crypto.Market{}, err
	}

	defer res.Body.Close()

	var ms marketResponse
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms {
		marketList, err = crypto.AppendMarket(marketList, m.Asset, m.PricedIn)
		if err != nil {
			log.Debugf("skipping market %s/%s: %v", m.Asset, m.PricedIn, err)
		}
	}

	return marketList, nil
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
		return ob, errors.Wrap(err, "unable to GET orderbook")
	}

	defer res.Body.Close()

	var r obResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return ob, errors.Wrap(err, "unable to decode response")
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
		err = errors.Wrapf(err, "unable to fetch Order Book")
	}

	return
}
