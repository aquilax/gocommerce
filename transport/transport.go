package transport

import (
	"io"
)

// Transport is general transport interface
type Transport interface {
	URL(path string, params map[string]string) (string, error)
	Post(url string, reader io.Reader) error
	Get(url string) ([]byte, error)
	Put(url string, reader io.Reader) error
	Delete(url string) error
}
