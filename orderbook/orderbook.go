package orderbook

import (
	"encoding/json"
	"sort"

	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

type (
	PendingOrder struct {
		Price,
		Volume float64
	}

	OrderBook struct {
		Bids,
		Asks []PendingOrder
	}

	ObResponse struct {
		Bids []interface{} `json:"bids"`
		Asks []interface{} `json:"asks"`
	}
)

func AsNestedSlice(os []PendingOrder) (orders [][]float64) {
	for _, o := range os {
		orders = append(orders, []float64{o.Price, o.Volume})
	}

	return
}

func calcVolume(orders []PendingOrder) (volume float64) {
	for _, o := range orders {
		volume += o.Volume
	}

	return
}

// 4. catches a case where BOTH sides of the orderbook are empty
func EnsureVolume(ob OrderBook) (OrderBook, error) {
	if calcVolume(ob.Asks) > 0 {
		return ob, nil
	}

	if calcVolume(ob.Bids) > 0 {
		return ob, nil
	}

	return OrderBook{}, errors.New("both sides of the order book are empty")
}

// 3. Goes through both sides of the order book and makes sure they're sorted by price
func Sort(ob OrderBook) (OrderBook, error) {
	sort.Slice(ob.Bids, func(i, j int) bool { return ob.Bids[i].Price > ob.Bids[j].Price })
	sort.Slice(ob.Asks, func(i, j int) bool { return ob.Asks[i].Price < ob.Asks[j].Price })

	return EnsureVolume(ob)
}

func tryPrice(order map[string]interface{}) (price interface{}) {
	price, ok := order["price"]
	if ok {
		return
	}

	price, ok = order["Rate"] // bittrex
	if ok {
		return
	}

	return
}

func tryVolume(order map[string]interface{}) (volume interface{}) {
	volume, ok := order["volume"]
	if ok {
		return
	}

	volume, ok = order["Quantity"] // bittrex
	if ok {
		return
	}

	return
}

func normaliseOrders(raw []interface{}) (orders []PendingOrder) {
	for i, rawOrder := range raw {
		var rawPrice, rawVol interface{}

		switch order := rawOrder.(type) {
		case []interface{}:
			rawPrice, rawVol = order[0], order[1]
			break

		case map[string]interface{}:
			rawPrice, rawVol = tryPrice(order), tryVolume(order)
			break

		default:
			panic(errors.Errorf("unknown order type returned %#v", order))
		}

		price, err := utils.ToFloat(rawPrice)
		if err != nil {
			panic(errors.Wrapf(err, "bids[%d].price = (%v) not convertible to float64", i, rawPrice))
		}

		vol, err := utils.ToFloat(rawVol)
		if err != nil {
			panic(errors.Wrapf(err, "bids[%d].volume = (%v) not convertible to float64", i, rawVol))
		}

		orders = append(orders, PendingOrder{price, vol})
	}

	return
}

// 2. takes two sides of an order book, and normalises them to []PendingOrder-s
func Normalise(asks, bids []interface{}) (ob OrderBook, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("unable to convert order book: %s", r)
		}
	}()

	return Sort(OrderBook{
		normaliseOrders(bids),
		normaliseOrders(asks),
	})

}

// 1. GETs a provided url and converts it to an orderbook
func Download(url string) (OrderBook, error) {
	res, err := utils.NetClient().Get(url)
	if err != nil {
		return OrderBook{}, err
	}

	defer res.Body.Close()

	var r ObResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return OrderBook{}, err
	}

	return Normalise(r.Asks, r.Bids)
}
