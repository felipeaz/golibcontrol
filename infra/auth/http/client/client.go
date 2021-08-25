package client

import (
	"net/http"
)

type HTTPClient struct{}

func NewHTTPClient() HTTPClient {
	return HTTPClient{}
}

func (c HTTPClient) Do(req *http.Request) (*http.Response, error) {
	cli := http.DefaultClient
	return cli.Do(req)
}
