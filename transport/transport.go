package transport

import (
	"io"
)

// Transport is general transport interface
type Transport interface {
	URL(path string, params map[string]string) (string, error)
	Get(url string) ([]byte, error)
	Post(url string, reader io.Reader) error
	Patch(url string, reader io.Reader) error
}
