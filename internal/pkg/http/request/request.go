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

func (h HttpRequest) PostRequest(id string, body []byte) ([]byte, error) {
	respBody := bytes.NewBuffer(body)
	resp, err := http.Post(h.Host+id, "application/json", respBody)
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

func (h HttpRequest) PostRequestWithoutBody(concatUrl string) ([]byte, error) {
	req, err := http.NewRequest("POST", h.Host+concatUrl, nil)
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

func (h HttpRequest) DeleteRequest(concatUrl string) error {
	req, err := http.NewRequest("DELETE", h.Host+concatUrl, nil)
	if err != nil {
		logger.LogError(err)
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}
