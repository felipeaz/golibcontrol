package _interface

type HTTPRequestInterface interface {
	Get(id string) ([]byte, error)
	Post(route string, body []byte) ([]byte, error)
	PostWithHeader(route string, body []byte, headerName, headerValue string) ([]byte, error)
	PostWithoutBody(route string) ([]byte, error)
	Delete(route string) error
}
