package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

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

type marketResponse []struct {
	Symbol string `json:"symbol"`
}

var (
	marketList []crypto.Market
	aliases    = []string{currencies.Bcc}
)

func Markets() (_ []crypto.Market, err error) {
	if len(marketList) == 0 {
		var res *http.Response
		res, err = utils.NetClient().Get(marketsUrl)
		if err != nil {
			return
		}

		defer res.Body.Close()

		var ms marketResponse
		err = json.NewDecoder(res.Body).Decode(&ms)
		if err != nil {
			return
		}

		var symbols []string
		for symbol := range currencies.All() {
			symbols = append(symbols, symbol)
		}

		orSymbols := strings.Join(symbols, "|")

		regexpStr := fmt.Sprintf("^(%s)(%s)$", orSymbols, orSymbols)

		regExp, err := regexp.Compile(regexpStr)
		if err != nil {
			return []crypto.Market{}, nil
		}

		for _, m := range ms {
			x := regExp.FindAllStringSubmatch(m.Symbol, -1)

			if len(x) == 1 && len(x[0]) == 3 {
				marketList = append(marketList, crypto.NewMarket(x[0][1], x[0][2]))
				continue
			}

			crypto.Log().Debugln("binance: at least one currency not recognised in", m.Symbol)
		}
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
		err = errors.Wrapf(err, "unable to fetch %s Order Book", Domain)
	}

	return
}
