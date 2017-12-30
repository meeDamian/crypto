package binance

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
	marketsUrl   = "https://api.binance.com/api/v1/ticker/allPrices"
	orderBookUrl = "https://api.binance.com/api/v1/depth?symbol=%s%s"
)

type marketRes []struct {
	Symbol string `json:"symbol"`
}

var marketList []crypto.Market

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) > 0 {
		return marketList, nil
	}

	res, err := utils.NetClient().Get(marketsUrl)
	if err != nil {
		return []crypto.Market{}, err
	}

	defer res.Body.Close()

	var ms marketRes
	err = json.NewDecoder(res.Body).Decode(&ms)
	if err != nil {
		return
	}

	for _, m := range ms {
		market, err := crypto.NewMarketFromSymbol(m.Symbol)
		if err != nil {
			log.Debugln("unable to parse symbol:", err)
			continue
		}

		marketList = append(marketList, market)
	}

	return marketList, nil
}

func OrderBook(m crypto.Market) (ob orderbook.OrderBook, err error) {
	url := fmt.Sprintf(orderBookUrl,
		currencies.Morph(m.Asset, aliases),
		currencies.Morph(m.PricedIn, aliases),
	)

	ob, err = orderbook.Download(url)
	if err != nil {
		err = errors.Wrap(err, "unable to fetch Order Book")
	}

	return
}
