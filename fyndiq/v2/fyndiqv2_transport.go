package fyndiqv2

import (
	"fmt"
	"github.com/aquilax/gocommerce/transport"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const BaseURL = "https://api.fyndiq.com/v2/"

type FyndiqV2Transport interface {
	Client() *http.Client
	NewRequest(method, urlStr string, body io.Reader) (*http.Request, error)
	transport.Transport
}

type Transport struct {
	user      string
	token     string
	userAgent string
	client    *http.Client
}

func NewTransport(user, token, userAgent string) *Transport {
	return &Transport{
		user,
		token,
		userAgent,
		&http.Client{},
	}
}

func (t *Transport) Client() *http.Client {
	return t.client
}

func (t *Transport) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	var err error
	var req *http.Request
	if req, err = http.NewRequest(method, urlStr, body); err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", t.userAgent)
	req.SetBasicAuth(t.user, t.token)
	return req, nil
}

func (t *Transport) URL(path string, params map[string]string) (string, error) {
	var u *url.URL
	var err error
	if u, err = url.Parse(BaseURL); err != nil {
		return "", err
	}
	u.Path += path
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (t *Transport) Get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = t.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	if resp, err = t.Client().Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (t *Transport) submit(method, url string, reader io.Reader) error {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = t.NewRequest(method, url, reader); err != nil {
		return err
	}
	if resp, err = t.client.Do(req); err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return nil
}

func (t *Transport) Patch(url string, reader io.Reader) error {
	return t.submit("PATCH", url, reader)
}

func (t *Transport) Post(url string, reader io.Reader) error {
	return t.submit("POST", url, reader)
}
