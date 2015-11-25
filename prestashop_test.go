package gocommerce

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPrestashop(t *testing.T) {
	Convey("Given new Prestashop", t, func() {
		ps := NewPrestashop()
		Convey("Prestashop is not null", func() {
			So(ps, ShouldNotBeNil)
		})
	})
}
