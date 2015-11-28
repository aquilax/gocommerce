package gocommerce

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PrestaShopTrasport is the default implementation of PrestaShopTransport
// if it doesn't suit your needs, you can consider providing your own implementation
type PrestaShopTrasport struct {
	apiURL string
	key    string
	client *http.Client
}

// NewPrestaShopTrasport creates new PrestaShopTrasport
func NewPrestaShopTrasport(apiURL, key string) *PrestaShopTrasport {
	return &PrestaShopTrasport{
		apiURL,
		key,
		&http.Client{},
	}
}

// getUrl builds URL for accessing PrestaShop
func (pt *PrestaShopTrasport) getURL(path string, params map[string]string) string {
	var err error
	var u *url.URL
	if u, err = url.Parse(pt.apiURL); err != nil {
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
func (pt *PrestaShopTrasport) Get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	req.SetBasicAuth(pt.key, "")
	if resp, err = pt.client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
