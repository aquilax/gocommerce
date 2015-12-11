package transport

import (
	"io"
)

type DummyTransport struct{}

func (dt *DummyTransport) URL(path string, params map[string]string) (string, error) {
	return "", nil
}

func (dt *DummyTransport) Get(url string) ([]byte, error) {
	return nil, nil
}

func (dt *DummyTransport) Post(url string, reader io.Reader) error {
	return nil
}

func (dt *DummyTransport) Put(url string, reader io.Reader) error {
	return nil
}

func (dt *DummyTransport) Patch(url string, reader io.Reader) error {
	return nil
}

func (dt *DummyTransport) Delete(url string) error {
	return nil
}
