package consumer

import (
	"encoding/json"
	"github.com/FelipeAz/golibcontrol/infra/http/interface"

	"github.com/FelipeAz/golibcontrol/infra/consumer/model"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type Consumer struct {
	HTTPRequest _interface.HTTPRequestInterface
	Log         logger.LogInterface
	JWTSecret   string
}

func NewConsumer(httpRequest _interface.HTTPRequestInterface, logger logger.LogInterface, jwtSecret string) Consumer {
	return Consumer{
		HTTPRequest: httpRequest,
		Log:         logger,
		JWTSecret:   jwtSecret,
	}
}

func (c Consumer) CreateConsumer(username string) (*model.Consumer, error) {
	body, _ := json.Marshal(model.Consumer{Username: username})

	b, err := c.HTTPRequest.Post("", body)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	var consumer *model.Consumer
	err = json.Unmarshal(b, &consumer)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return consumer, nil
}

func (c Consumer) CreateConsumerKey(consumerId string) (*model.ConsumerKey, error) {
	concatUrl := consumerId + "/jwt"
	err := c.HTTPRequest.Delete(concatUrl)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	body := model.CreateKeyBody{Secret: c.JWTSecret}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	b, err := c.HTTPRequest.Post(concatUrl, bodyBytes)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	var consumerKey *model.ConsumerKey
	err = json.Unmarshal(b, &consumerKey)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return consumerKey, nil
}

func (c Consumer) GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, error) {
	concatUrl := consumerId + "/jwt/" + keyId

	b, err := c.HTTPRequest.Get(concatUrl)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	var consumerKey *model.ConsumerKey
	err = json.Unmarshal(b, &consumerKey)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return consumerKey, nil
}

func (c Consumer) GetAllConsumerKeys(consumerId string) (*model.Keys, error) {
	concatUrl := consumerId + "/jwt/"

	b, err := c.HTTPRequest.Get(concatUrl)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	var consumerKeys *model.Keys
	err = json.Unmarshal(b, &consumerKeys)
	if err != nil {
		c.Log.Error(err)
		return nil, err
	}

	return consumerKeys, nil
}

func (c Consumer) RetrieveConsumerKey(consumerId string) (*model.ConsumerKey, error) {
	var consumerKey *model.ConsumerKey
	keys, apiError := c.GetAllConsumerKeys(consumerId)
	if apiError != nil {
		return nil, apiError
	}

	if keys == nil || len(keys.Data) == 0 {
		consumerKey, apiError = c.CreateConsumerKey(consumerId)
		if apiError != nil {
			return nil, apiError
		}
		return consumerKey, nil
	}

	return &keys.Data[0], nil
}

func (c Consumer) DeleteConsumerKey(consumerId, consumerKeyId string) error {
	concatUrl := consumerId + "/jwt/" + consumerKeyId
	err := c.HTTPRequest.Delete(concatUrl)
	if err != nil {
		c.Log.Error(err)
		return err
	}
	return nil
}

func (c Consumer) DeleteConsumer(consumerId string) error {
	err := c.HTTPRequest.Delete(consumerId)
	if err != nil {
		c.Log.Error(err)
		return err
	}
	return nil
}
