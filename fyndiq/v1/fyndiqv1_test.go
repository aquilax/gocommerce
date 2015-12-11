package fyndiqv1

import (
	"github.com/aquilax/gocommerce/transport"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"testing"
)

type FV1DummyTransport struct {
	transport.DummyTransport
}

func (t *FV1DummyTransport) Client() *http.Client {
	return nil
}

func (t *FV1DummyTransport) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	return nil, nil
}

func TestFyndiqV2(t *testing.T) {
	Convey("Given Transport", t, func() {
		t := NewTransport("user", "token")
		Convey("Given new FyndiqV1", func() {
			f := New(t)
			Convey("FyndiqV1 is not null", func() {
				So(f, ShouldNotBeNil)
			})
		})
	})
}
