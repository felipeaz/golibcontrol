package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelipeAz/golibcontrol/infra/auth/http/client"
	"github.com/FelipeAz/golibcontrol/infra/auth/http/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateConsumerSuccess(t *testing.T) {
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

func TestCreateConsumerUnmarshalError(t *testing.T) {
	username := "email@test.com"
	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(nil))
			}
		}),
	)
	defer testServer.Close()
	cli := request.NewHttpRequest(client.NewHTTPClient(), testServer.URL+"/")
	auth := NewAuth(cli)

	consumer, err := auth.CreateConsumer(username)

	assert.Nil(t, consumer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unexpected end of JSON input")
}

func TestCreateConsumerKeySuccess(t *testing.T) {
	consumerId := "49eafa57-d530-4ddc-a399-7df4a30225d2"
	secret := "98bf1013-b69f-430b-b4f4-822a9c4e3d59"
	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == consumerId+"/jwt" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{"id": "49eafa57-d530-4ddc-a399-7df4a30225d2", "custom_id": "123123", "username": "email@test.com"}`))
			}
		}),
	)
	defer testServer.Close()
	cli := request.NewHttpRequest(client.NewHTTPClient(), testServer.URL+"/")
	auth := NewAuth(cli)

	consumerKey, err := auth.CreateConsumerKey(consumerId, secret)

	assert.NoError(t, err)
	assert.NotEmpty(t, consumerKey.Id)
	assert.NotEmpty(t, consumerKey.Consumer.Id)
}
