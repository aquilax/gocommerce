package gocommerce

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFyndiqV2(t *testing.T) {
	Convey("Given Transport", t, func() {
		tt := &TestTransport{}
		Convey("Given new FyndiqV2", func() {
			f := NewFyndiqV2(tt)
			Convey("FyndiqV2 is not null", func() {
				So(f, ShouldNotBeNil)
			})
		})
	})
}
