package fyndiqv2

import (
	"github.com/aquilax/gocommerce/transport"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"testing"
)

type FV2DummyTransport struct {
	transport.DummyTransport
}

func (t *FV2DummyTransport) Client() *http.Client {
	return nil
}

func (t *FV2DummyTransport) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	return nil, nil
}

func TestFyndiqV2(t *testing.T) {
	Convey("Given Transport", t, func() {
		tt := &FV2DummyTransport{}
		Convey("Given new FyndiqV2", func() {
			f := New(tt)
			Convey("FyndiqV2 is not null", func() {
				So(f, ShouldNotBeNil)
			})
		})
	})
}
