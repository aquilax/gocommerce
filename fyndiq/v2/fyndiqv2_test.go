package fyndiqv2

import (
	"github.com/aquilax/gocommerce/transport"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFyndiqV2(t *testing.T) {
	Convey("Given Transport", t, func() {
		tt := &transport.DummyTransport{}
		Convey("Given new FyndiqV2", func() {
			f := New(tt)
			Convey("FyndiqV2 is not null", func() {
				So(f, ShouldNotBeNil)
			})
		})
	})
}
