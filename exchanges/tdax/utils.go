package tdax

import (
	"encoding/json"
	"github.com/meeDamian/crypto/markets"
	"math"

	"github.com/meeDamian/crypto/currencies"
	"github.com/meeDamian/crypto/orderbook"
	"github.com/meeDamian/crypto/utils"
	"github.com/pkg/errors"
)

const currenciesUrl = "https://api.tdax.com/public/getcurrencies"

type currencyResponse []struct {
	Name    string `json:"Currency"`
	Divider int64  `json:"Divider"`
}

var precisions = make(map[string]int)

func currencyPrecisions() (err error) {
	if len(precisions) > 0 {
		return
	}

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

func normalisedPendingOrder(o order, m markets.Market) (po orderbook.PendingOrder, err error) {
	volume, err := normalize(m.Asset, o.Volume)
	if err != nil {
		return po, err
	}

	price, err := normalize(m.PricedIn, o.Price)
	if err != nil {
		return po, err
	}

	return orderbook.PendingOrder{
		Price:  price,
		Volume: volume,
	}, nil
}

func normalize(currency string, amount float64) (float64, error) {
	a, err := currencies.Get(currency)
	if err != nil {
		return 0, err
	}

	precision, ok := precisions[a.Name]
	if !ok {
		return 0, errors.Errorf("precision of %s unknown", currency)
	}

	return amount / math.Pow10(precision), nil
}
