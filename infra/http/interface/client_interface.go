package _interface

import (
	"net/http"
)

type HTTPClientInterface interface {
	Do(r *http.Request) (*http.Response, error)
}
