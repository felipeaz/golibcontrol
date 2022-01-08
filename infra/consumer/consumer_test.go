package consumer

import (
	"encoding/json"
	"github.com/FelipeAz/golibcontrol/infra/http/client"
	"github.com/FelipeAz/golibcontrol/infra/http/request"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelipeAz/golibcontrol/infra/consumer/model"
	"github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/stretchr/testify/assert"
)

func TestCreateConsumerSuccess(t *testing.T) {
	username := "email@test.com"
	expectedConsumer := &model.Consumer{
		Id:        "49eafa57-d530-4ddc-a399-7df4a30225d2",
		CustomId:  "123123",
		Username:  username,
		Tags:      nil,
		CreatedAt: 0,
	}
	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.Header().Add("Content-Type", "application/json")
				resp, err := json.Marshal(expectedConsumer)
				if err != nil {
					assert.Fail(t, "Failed to marshal expected response")
				}
				if _, err := w.Write(resp); err != nil {
					assert.Fail(t, "Failed to Write test response")
				}
			}
		}),
	)
	defer testServer.Close()
	cli := request.NewHttpRequest(client.NewHTTPClient(), testServer.URL+"/")
	loggerMock := new(logger.Mock)
	jwtSecret := "jwtsecret"
	auth := NewConsumer(cli, loggerMock, jwtSecret)

	consumer, err := auth.CreateConsumer(username)

	assert.NoError(t, err)
	assert.Equal(t, consumer, expectedConsumer)
}

func TestCreateConsumerUnmarshalError(t *testing.T) {
	username := "email@test.com"
	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(nil))
				if err != nil {
					t.Fail()
				}
			}
		}),
	)
	defer testServer.Close()
	cli := request.NewHttpRequest(client.NewHTTPClient(), testServer.URL+"/")
	loggerMock := new(logger.Mock)

	jwtSecret := "jwtsecret"
	auth := NewConsumer(cli, loggerMock, jwtSecret)

	consumer, err := auth.CreateConsumer(username)

	assert.Nil(t, consumer)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unexpected end of JSON input")
}

func TestCreateConsumerHTTPRequestError(t *testing.T) {
	username := "email@test.com"

	cli := request.NewHttpRequest(client.NewHTTPClient(), "")
	loggerMock := new(logger.Mock)
	jwtSecret := "jwtsecret"
	auth := NewConsumer(cli, loggerMock, jwtSecret)

	consumer, err := auth.CreateConsumer(username)
	assert.NotNil(t, err)
	assert.Nil(t, consumer)
}

func TestCreateConsumerKeySuccess(t *testing.T) {
	consumerId := "49eafa57-d530-4ddc-a399-7df4a30225d2"
	expectedConsumer := &model.ConsumerKey{
		Key:       "",
		CreatedAt: 0,
		Id:        "",
		Consumer: struct {
			Id string `json:"id"`
		}{},
		Tags:         nil,
		RsaPublicKey: nil,
		Algorithm:    "",
		Secret:       "",
	}
	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/"+consumerId+"/jwt" {
				w.Header().Add("Content-Type", "application/json")
				resp, err := json.Marshal(expectedConsumer)
				if err != nil {
					assert.Fail(t, "Failed to marshal expected response")
				}
				if _, err := w.Write(resp); err != nil {
					assert.Fail(t, "Failed to Write test response")
				}
			}
		}),
	)
	defer testServer.Close()
	cli := request.NewHttpRequest(client.NewHTTPClient(), testServer.URL+"/")
	loggerMock := new(logger.Mock)
	jwtSecret := "jwtsecret"
	auth := NewConsumer(cli, loggerMock, jwtSecret)

	consumerKey, err := auth.CreateConsumerKey(consumerId)

	assert.NoError(t, err)
	assert.Equal(t, consumerKey, expectedConsumer)
}
