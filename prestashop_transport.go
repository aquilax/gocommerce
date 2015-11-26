package gocommerce

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PrestaShopTransport is transport interface for accessing PrestaShop API
type PrestaShopTransport interface {
	Get(url string) ([]byte, error)
}

// DefaultPrestaShopTrasport is the default implementation of PrestaShopTransport
// if it doesn't suit your needs, you can consider providing your own implementation
type DefaultPrestaShopTrasport struct {
	apiURL string
	key    string
	client *http.Client
}

// NewDefaultPrestaShopTrasport creates new DefaultPrestaShopTrasport
func NewDefaultPrestaShopTrasport(apiURL, key string) *DefaultPrestaShopTrasport {
	return &DefaultPrestaShopTrasport{
		apiURL,
		key,
		&http.Client{},
	}
}

// getUrl builds URL for accessing PrestaShop
func (dpt *DefaultPrestaShopTrasport) getURL(path string, params map[string]string) string {
	var err error
	var u *url.URL
	if u, err = url.Parse(dpt.apiURL); err != nil {
		panic(err)
	}
	u.Path += path
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// Get sends GET request to PrestaShop
func (dpt *DefaultPrestaShopTrasport) Get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	req.SetBasicAuth(dpt.key, "")
	if resp, err = dpt.client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
