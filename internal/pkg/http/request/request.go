package request

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/logger"
)

type HttpRequest struct {
	Host string
}

func NewHttpRequest(host string) HttpRequest {
	return HttpRequest{
		Host: host,
	}
}

func (h HttpRequest) GetRequest(id string) ([]byte, error) {
	resp, err := http.Get(h.Host + id)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) PostRequest(route string, body []byte) ([]byte, error) {
	respBody := bytes.NewBuffer(body)
	resp, err := http.Post(h.Host+route, "application/json", respBody)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) PostRequestWithHeader(route string, body []byte, headerName, headerValue string) ([]byte, error) {
	respBody := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", h.Host+route, respBody)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(headerName, headerValue)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) PostRequestWithoutBody(route string) ([]byte, error) {
	req, err := http.NewRequest("POST", h.Host+route, nil)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) DeleteRequest(route string) error {
	req, err := http.NewRequest("DELETE", h.Host+route, nil)
	if err != nil {
		logger.LogError(err)
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}
