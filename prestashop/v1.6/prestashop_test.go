package gocommerce

import (
	"github.com/aquilax/gocommerce/transport"

	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPrestashop(t *testing.T) {
	Convey("Given PrestashopTransport", t, func() {
		t := &transport.DummyTransport{}
		Convey("Given new PrestaShop", func() {
			api := New(t)
			Convey("PrestaShop is not null", func() {
				So(api, ShouldNotBeNil)
			})
		})
	})
}
