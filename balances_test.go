package crypto

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBalances_Add_UnknownCurrency(t *testing.T) {
	Convey("should return error on unknown currency", t, func() {
		balances := make(Balances)
		err := balances.Add("XXX", nil, nil, nil)

		So(err, ShouldNotBeNil)
	})
}

func TestBalances_Add_NilAvailable(t *testing.T) {
	Convey("should return error if 'available' missing", t, func() {
		balances := make(Balances)
		err := balances.Add("USD", nil, nil, nil)

		So(err, ShouldNotBeNil)
	})
}

func TestBalances_Add_BrokenStringAvailable(t *testing.T) {
	Convey("should return error if 'available' cannot be parsed to float64", t, func() {
		balances := make(Balances)
		err := balances.Add("USD", "NOT_A_NUMBER", nil, nil)

		So(err, ShouldNotBeNil)
	})
}

func TestBalances_Add_UnknownTypeAvailable(t *testing.T) {
	Convey("should return error if 'available' cannot be parsed to float64", t, func() {
		balances := make(Balances)
		err := balances.Add("USD", []string{}, nil, nil)

		So(err, ShouldNotBeNil)
	})
}
