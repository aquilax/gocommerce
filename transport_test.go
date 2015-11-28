package gocommerce

import (
	"io"
)

type TestTransport struct{}

func (tt *TestTransport) URL(path string, params map[string]string) (string, error) {
	return "", nil
}

func (tt *TestTransport) Get(url string) ([]byte, error) {
	return nil, nil
}

func (tt *TestTransport) Patch(url string, reader io.Reader) error {
	return nil
}
