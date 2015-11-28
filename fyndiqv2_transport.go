package gocommerce

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const FyndiqV2BaseURL = "https://api.fyndiq.com/v2/"

type FyndiqV2Trasport struct {
	user   string
	token  string
	client *http.Client
}

func NewFyndiqV2Trasport(user, token string) *FyndiqV2Trasport {
	return &FyndiqV2Trasport{
		user,
		token,
		&http.Client{},
	}
}

func (ft *FyndiqV2Trasport) getURL(path string, params map[string]string) string {
	var err error
	var u *url.URL
	if u, err = url.Parse(FyndiqV2BaseURL); err != nil {
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

func (ft *FyndiqV2Trasport) Get(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	req.SetBasicAuth(ft.user, ft.token)
	if resp, err = ft.client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
