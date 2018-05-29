package tdax

import (
	"encoding/json"
	"fmt"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const (
	orderBookUrl  = "https://api.tdax.com/orders?Symbol=%s_%s"
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
)

var (
	marketList []crypto.Market

	aliases = []string{currencies.Rpx}
)

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
		po, err := normalisedPendingOrder(o, m)
		if err != nil {
			return ob, err
		}

		asks = append(asks, po)
	}

	for _, o := range r.Bids {
		po, err := normalisedPendingOrder(o, m)
		if err != nil {
			return ob, err
		}

		bids = append(bids, po)
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
