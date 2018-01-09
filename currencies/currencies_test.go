package currencies

import (
	"testing"

	"github.com/meeDamian/arb/lib/currencies"
	. "github.com/meeDamian/crypto/currencies/symbols"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMorphBccMatch(t *testing.T) {
	a := Morph(Bch, []string{currencies.Bcc})

	Convey("Should change BCH to BCC", t, func() {
		So(a, ShouldEqual, Bcc)
	})
}

func TestMorphBccNoMatch(t *testing.T) {
	a := Morph(Btc, []string{currencies.Bcc})

	Convey("Should change BCH to BCC", t, func() {
		So(a, ShouldEqual, Btc)
	})
}

func TestMorphXbtMatch(t *testing.T) {
	a := Morph(Btc, []string{currencies.Xbt})

	Convey("Should change BCH to BCC", t, func() {
		So(a, ShouldEqual, Xbt)
	})
}

func TestMorphXbtNoMatch(t *testing.T) {
	a := Morph(Eth, []string{currencies.Xbt})

	Convey("Should change BCH to BCC", t, func() {
		So(a, ShouldEqual, Eth)
	})
}

func TestGetFiat(t *testing.T) {
	Convey("Should return USD object", t, func() {
		usd, err := Get(Usd)
		So(err, ShouldBeNil)
		So(usd, ShouldNotBeNil)
		So(usd, ShouldResemble, list[Usd])
	})
}

func TestGetCrypto(t *testing.T) {
	Convey("Should return BTC object", t, func() {
		usd, err := Get(Btc)
		So(err, ShouldBeNil)
		So(usd, ShouldNotBeNil)
		So(usd, ShouldResemble, list[Btc])
	})
}

func TestGetNonexistent(t *testing.T) {
	Convey("Should error", t, func() {
		usd, err := Get("AAA")
		So(err, ShouldNotBeNil)
		So(usd, ShouldBeZeroValue)
	})
}
