package gocommerce

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestPrestaShopTransport struct{}

func (tpt *TestPrestaShopTransport) get(url string) ([]byte, error) {
	return nil, nil
}

func TestPrestashop(t *testing.T) {
	Convey("Given PrestashopTransport", t, func() {
		tpt := &TestPrestaShopTransport{}
		Convey("Given new PrestaShop", func() {
			ps := NewPrestaShop(tpt)
			Convey("PrestaShop is not null", func() {
				So(ps, ShouldNotBeNil)
			})
		})
	})
}
