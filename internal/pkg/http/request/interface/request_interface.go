package _interface

type HTTPRequestInterface interface {
	GetRequest(id string) ([]byte, error)
	PostRequest(route string, body []byte) ([]byte, error)
	PostRequestWithHeader(route string, body []byte, headerName, headerValue string) ([]byte, error)
	PostRequestWithoutBody(route string) ([]byte, error)
	DeleteRequest(route string) error
}
