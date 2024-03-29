package request

import (
	"bytes"
	"github.com/FelipeAz/golibcontrol/infra/http/interface"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpRequest struct {
	Client  _interface.HTTPClientInterface
	BaseUrl string
}

func NewHttpRequest(client _interface.HTTPClientInterface, baseUrl string) HttpRequest {
	return HttpRequest{
		Client:  client,
		BaseUrl: baseUrl,
	}
}

func (h HttpRequest) Get(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", h.BaseUrl+id, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) Post(route string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", h.BaseUrl+route, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := h.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) PostWithHeader(route string, body []byte, headerName, headerValue string) ([]byte, error) {
	req, err := http.NewRequest("POST", h.BaseUrl+route, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(headerName, headerValue)
	resp, err := h.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) PostWithoutBody(route string) ([]byte, error) {
	req, err := http.NewRequest("POST", h.BaseUrl+route, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return b, nil
}

func (h HttpRequest) Delete(route string) error {
	req, err := http.NewRequest("DELETE", h.BaseUrl+route, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}
