package utils

import (
	"strconv"
	"testing"

	"github.com/meeDamian/crypto/currencies"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFloat64ToFloat(t *testing.T) {
	Convey("Should just pipe a float through", t, func() {
		out, err := ToFloat(666.666)
		So(err, ShouldBeNil)
		So(out, ShouldEqual, 666.666)
	})
}

func TestStringToFloat(t *testing.T) {
	Convey("Should convert string to float", t, func() {
		out, err := ToFloat("777.777")
		So(err, ShouldBeNil)
		So(out, ShouldEqual, 777.777)
	})
}

func TestInvalidStringToFloat(t *testing.T) {
	Convey("Should panic for invalid string", t, func() {
		_, err := ToFloat("NOT A NUMBER")
		So(err, ShouldBeError)
		So(err, ShouldHaveSameTypeAs, &strconv.NumError{})
		So(err.Error(), ShouldContainSubstring, "parsing")
	})
}

func TestOtherToFloat(t *testing.T) {
	Convey("Should panic for nil", t, func() {
		_, err := ToFloat(nil)
		So(err, ShouldBeError)
		So(err.Error(), ShouldContainSubstring, "unsupported type")
	})

	Convey("Should panic for bool", t, func() {
		_, err := ToFloat(true)
		So(err, ShouldBeError)
		So(err.Error(), ShouldContainSubstring, "unsupported type")
	})

	Convey("Should panic for anon struct", t, func() {
		_, err := ToFloat(struct{}{})
		So(err, ShouldBeError)
		So(err.Error(), ShouldContainSubstring, "unsupported type")
	})

	Convey("Should panic for any object", t, func() {
		_, err := ToFloat(currencies.Currency{})
		So(err, ShouldBeError)
		So(err.Error(), ShouldContainSubstring, "unsupported type")
	})
}
