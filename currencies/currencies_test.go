package currencies

import (
	. "github.com/meeDamian/crypto/currencies/symbols"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMorphBccMatch(t *testing.T) {
	a := MorphBcc(Bch)

	Convey("Should change BCH to BCC", t, func() {
		So(a, ShouldEqual, Bcc)
	})
}

func TestMorphBccNoMatch(t *testing.T) {
	a := MorphBcc(Btc)

	Convey("Should change BCH to BCC", t, func() {
		So(a, ShouldEqual, Btc)
	})
}

func TestMorphXbtMatch(t *testing.T) {
	a := MorphXbt(Btc)

	Convey("Should change BCH to BCC", t, func() {
		So(a, ShouldEqual, Xbt)
	})
}

func TestMorphXbtNoMatch(t *testing.T) {
	a := MorphXbt(Eth)

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
		So(err, ShouldBeNil)
		So(usd, ShouldBeZeroValue)
	})
}
