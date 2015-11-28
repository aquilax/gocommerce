package gocommerce

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPrestashop(t *testing.T) {
	Convey("Given PrestashopTransport", t, func() {
		tt := &TestTransport{}
		Convey("Given new PrestaShop", func() {
			ps := NewPrestaShop(tt)
			Convey("PrestaShop is not null", func() {
				So(ps, ShouldNotBeNil)
			})
		})
	})
}
