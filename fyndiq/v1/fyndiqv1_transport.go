package fyndiqv1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const BaseURL = "https://fyndiq.se/api/v1"

type Transport struct {
	user   string
	token  string
	client *http.Client
}

func NewTransport(user string, token string) *Transport {
	return &Transport{
		user,
		token,
		&http.Client{},
	}
}

func (t *Transport) URL(path string, params map[string]string) (string, error) {
	var u *url.URL
	var err error
	if u, err = url.Parse(BaseURL); err != nil {
		return "", err
	}
	u.Path += path
	q := u.Query()
	// add auth
	q.Set("user", t.user)
	q.Set("token", t.token)
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (t *Transport) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	var err error
	var req *http.Request
	if req, err = http.NewRequest(method, urlStr, body); err != nil {
		return nil, err
	}
	return req, nil
}

func (t *Transport) Get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	if resp, err = t.client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (t *Transport) Delete(url string) error {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("DELETE", url, nil); err != nil {
		return err
	}
	if resp, err = t.client.Do(req); err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return nil
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

func (t *Transport) Post(url string, reader io.Reader) error {
	return t.submit("POST", url, reader)
}

func (t *Transport) Put(url string, reader io.Reader) error {
	return t.submit("PUT", url, reader)
}
