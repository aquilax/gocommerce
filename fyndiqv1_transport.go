package gocommerce

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const FyndiqV1BaseURL = "https://fyndiq.se/api/v1"

type FyndiqV1Transport struct {
	user   string
	token  string
	client *http.Client
}

func NewFyndiqV1Transport(user string, token string) *FyndiqV1Transport {
	return &FyndiqV1Transport{
		user,
		token,
		&http.Client{},
	}
}

func (dft *FyndiqV1Transport) getURL(path string, params map[string]string) string {
	var err error
	var u *url.URL
	if u, err = url.Parse(FyndiqV1BaseURL); err != nil {
		panic(err)
	}
	u.Path += path
	q := u.Query()
	// add auth
	q.Set("user", dft.user)
	q.Set("token", dft.token)
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (dft *FyndiqV1Transport) Get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	if resp, err = dft.client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
