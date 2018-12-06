package crypto

import (
	"github.com/meeDamian/crypto/markets"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseSymbolKraken(t *testing.T) {
	Convey("should be able to process Kraken symbols", t, func() {
		market, err := markets.NewFromSymbol("XXBTZEUR")

		So(err, ShouldBeNil)
		So(market.Asset, ShouldEqual, "BTC")
		So(market.PricedIn, ShouldEqual, "EUR")
	})
}

func TestParseSymbolNormal(t *testing.T) {
	Convey("should be able to process normal symbols", t, func() {
		market, err := markets.NewFromSymbol("BTCUSD")

		So(err, ShouldBeNil)
		So(market.Asset, ShouldEqual, "BTC")
		So(market.PricedIn, ShouldEqual, "USD")
	})
}

func TestParseSymbolSlash(t *testing.T) {
	Convey("should be able to process symbols with a slash", t, func() {
		market, err := markets.NewFromSymbol("ETH/BTC")

		So(err, ShouldBeNil)
		So(market.Asset, ShouldEqual, "ETH")
		So(market.PricedIn, ShouldEqual, "BTC")
	})
}

func TestParseSymbolAliases(t *testing.T) {
	Convey("should be able to process symbols using aliases", t, func() {
		market, err := markets.NewFromSymbol("DRK/XBT")

		So(err, ShouldBeNil)
		So(market.Asset, ShouldEqual, "DASH")
		So(market.PricedIn, ShouldEqual, "BTC")
	})
}

func TestParseSymbolUnknownCurrency(t *testing.T) {
	Convey("should be able to process symbols using aliases", t, func() {
		_, err := markets.NewFromSymbol("XXXBTC")

		So(err, ShouldNotBeNil)
	})
}
