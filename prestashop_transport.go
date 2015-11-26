package gocommerce

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type PrestaShopTransport interface {
	get(url string) ([]byte, error)
}

type DefaultPrestaShopTrasport struct {
	storeUrl string
	key      string
	client   *http.Client
}

func NewDefaultPrestaShopTrasport(storeURL, key string) *DefaultPrestaShopTrasport {
	return &DefaultPrestaShopTrasport{
		storeURL,
		key,
		&http.Client{},
	}
}

func (dpt *DefaultPrestaShopTrasport) get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	if resp, err = dpt.client.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK || resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
