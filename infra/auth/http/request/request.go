package request

import (
	"bytes"
	"io/ioutil"
	"net/http"

	_interface "github.com/FelipeAz/golibcontrol/infra/auth/http/interface"
	"github.com/FelipeAz/golibcontrol/infra/logger"
)

type HttpRequest struct {
	Client _interface.HTTPClientInterface
	Host   string
}

func NewHttpRequest(client _interface.HTTPClientInterface, host string) HttpRequest {
	return HttpRequest{
		Client: client,
		Host:   host,
	}
}

func (h HttpRequest) Get(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", h.Host+id, nil)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	resp, err := h.Client.Do(req)
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

func (h HttpRequest) Post(route string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", h.Host+route, bytes.NewBuffer(body))
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := h.Client.Do(req)
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

func (h HttpRequest) PostWithHeader(route string, body []byte, headerName, headerValue string) ([]byte, error) {
	req, err := http.NewRequest("POST", h.Host+route, bytes.NewBuffer(body))
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(headerName, headerValue)
	resp, err := h.Client.Do(req)
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

func (h HttpRequest) PostWithoutBody(route string) ([]byte, error) {
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

func (h HttpRequest) Delete(route string) error {
	req, err := http.NewRequest("DELETE", h.Host+route, nil)
	if err != nil {
		logger.LogError(err)
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}
