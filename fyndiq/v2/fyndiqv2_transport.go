package fyndiqv2

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const BaseURL = "https://api.fyndiq.com/v2/"

type Trasport struct {
	user   string
	token  string
	client *http.Client
}

func NewTrasport(user, token string) *Trasport {
	return &Trasport{
		user,
		token,
		&http.Client{},
	}
}

func (t *Trasport) URL(path string, params map[string]string) (string, error) {
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

func (t *Trasport) Get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	req.SetBasicAuth(t.user, t.token)
	if resp, err = t.client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (t *Trasport) Patch(url string, reader io.Reader) error {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("PATCH", url, reader); err != nil {
		return err
	}
	req.SetBasicAuth(t.user, t.token)
	if resp, err = t.client.Do(req); err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return nil
}
