package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelipeAz/golibcontrol/infra/auth/http/client"
	"github.com/FelipeAz/golibcontrol/infra/auth/http/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateConsumer(t *testing.T) {
	username := "email@test.com"
	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{"id": "49eafa57-d530-4ddc-a399-7df4a30225d2", "custom_id": "123123", "username": "email@test.com"}`))
			}
		}),
	)
	defer testServer.Close()
	cli := request.NewHttpRequest(client.NewHTTPClient(), testServer.URL+"/")
	auth := NewAuth(cli)

	consumer, err := auth.CreateConsumer(username)

	assert.NoError(t, err)
	assert.Equal(t, consumer.Username, username)
	assert.NotEmpty(t, consumer.Id)
	assert.NotEmpty(t, consumer.CustomId)
}
