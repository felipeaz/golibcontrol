package auth

import (
	"encoding/json"

	"github.com/FelipeAz/golibcontrol/infra/auth/model"
	"github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/internal/pkg/http/request/interface"
)

type Auth struct {
	HttpRequest _interface.HTTPRequestInterface
}

func NewAuth(httpRequest _interface.HTTPRequestInterface) Auth {
	return Auth{
		HttpRequest: httpRequest,
	}
}

func (a Auth) GetConsumer(consumerId string) (string, error) {
	b, err := a.HttpRequest.GetRequest(consumerId)
	if err != nil {
		logger.LogError(err)
		return "", err
	}
	var consumer model.Consumer
	err = json.Unmarshal(b, &consumer)
	if err != nil {
		logger.LogError(err)
		return "", err
	}

	return consumer.Id, nil
}

func (a Auth) CreateConsumer(username string) (*model.Consumer, error) {
	var consumer *model.Consumer
	body, err := json.Marshal(model.Consumer{Username: username})
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	b, err := a.HttpRequest.PostRequest("", body)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	err = json.Unmarshal(b, &consumer)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	if consumer == nil {
		return nil, err
	}

	return consumer, nil
}

func (a Auth) CreateConsumerKey(consumerId, secret string) (*model.ConsumerKey, error) {
	var consumerKey *model.ConsumerKey
	concatUrl := consumerId + "/jwt"

	err := a.HttpRequest.DeleteRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	body := model.CreateKeyBody{Secret: secret}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	b, err := a.HttpRequest.PostRequest(concatUrl, bodyBytes)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	err = json.Unmarshal(b, &consumerKey)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	if consumerKey == nil {
		return nil, err
	}

	return consumerKey, nil
}

func (a Auth) GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, error) {
	concatUrl := consumerId + "/jwt/" + keyId

	b, err := a.HttpRequest.GetRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	var consumerKey *model.ConsumerKey
	err = json.Unmarshal(b, &consumerKey)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return consumerKey, nil
}

func (a Auth) GetAllConsumerKeys(consumerId string) (*model.Keys, error) {
	concatUrl := consumerId + "/jwt/"

	b, err := a.HttpRequest.GetRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	var consumerKeys *model.Keys
	err = json.Unmarshal(b, &consumerKeys)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return consumerKeys, nil
}

func (a Auth) RetrieveConsumerKey(consumerId, secret string) (*model.ConsumerKey, error) {
	var consumerKey *model.ConsumerKey
	keys, apiError := a.GetAllConsumerKeys(consumerId)
	if apiError != nil {
		return nil, apiError
	}

	if keys == nil || len(keys.Data) == 0 {
		consumerKey, apiError = a.CreateConsumerKey(consumerId, secret)
		if apiError != nil {
			return nil, apiError
		}
		return consumerKey, nil
	}

	return &keys.Data[0], nil
}

func (a Auth) DeleteConsumer(consumerId, consumerKeyId string) error {
	concatUrl := consumerId + "/jwt/" + consumerKeyId
	err := a.HttpRequest.DeleteRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
