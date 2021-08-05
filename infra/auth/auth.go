package auth

import (
	"encoding/json"
	"net/http"

	authErrors "github.com/FelipeAz/golibcontrol/infra/auth/errors"
	"github.com/FelipeAz/golibcontrol/infra/auth/model"
	"github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/pkg/http/request"
)

type Auth struct {
	HttpRequest request.HttpRequest
}

func NewAuth(httpRequest request.HttpRequest) Auth {
	return Auth{
		HttpRequest: httpRequest,
	}
}

func (a Auth) GetConsumer(consumerId string) (string, *errors.ApiError) {
	b, err := a.HttpRequest.GetRequest(consumerId)
	if err != nil {
		logger.LogError(err)
		return "", &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToRetrieveConsumer,
			Error:   err.Error(),
		}
	}
	var consumer model.Consumer
	err = json.Unmarshal(b, &consumer)
	if err != nil {
		logger.LogError(err)
		return "", &errors.ApiError{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
	}

	return consumer.Id, nil
}

func (a Auth) CreateConsumer(username string) (*model.Consumer, *errors.ApiError) {
	var consumer *model.Consumer
	body, err := json.Marshal(model.Consumer{Username: username})
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToMarshalConsumer,
			Error:   err.Error(),
		}
	}

	b, err := a.HttpRequest.PostRequest("", body)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToRetrieveConsumer,
			Error:   err.Error(),
		}
	}

	err = json.Unmarshal(b, &consumer)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToUnmarshalConsumer,
			Error:   err.Error(),
		}
	}

	if consumer == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToCreateConsumer,
		}
	}

	return consumer, nil
}

func (a Auth) CreateConsumerKey(consumerId string) (*model.ConsumerKey, *errors.ApiError) {
	var consumerKey *model.ConsumerKey
	concatUrl := consumerId + "/key-auth/"

	err := a.HttpRequest.DeleteRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToCleanConsumerKeys,
			Error:   err.Error(),
		}
	}

	b, err := a.HttpRequest.PostRequestWithoutBody(concatUrl)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToRetrieveConsumer,
			Error:   err.Error(),
		}
	}

	err = json.Unmarshal(b, &consumerKey)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToUnmarshalConsumer,
			Error:   err.Error(),
		}
	}

	if consumerKey == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToCreateConsumer,
		}
	}
	return consumerKey, nil
}

func (a Auth) GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, *errors.ApiError) {
	concatUrl := consumerId + "/key-auth/" + keyId

	b, err := a.HttpRequest.GetRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToRetriveConsumerKeys,
			Error:   err.Error(),
		}
	}

	var consumerKey *model.ConsumerKey
	err = json.Unmarshal(b, &consumerKey)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToUnmarshalConsumer,
			Error:   err.Error(),
		}
	}

	return consumerKey, nil
}

func (a Auth) GetAllConsumerKeys(consumerId string) (*model.Keys, *errors.ApiError) {
	concatUrl := consumerId + "/key-auth/"

	b, err := a.HttpRequest.GetRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToRetriveConsumerKeys,
			Error:   err.Error(),
		}
	}

	var consumerKeys *model.Keys
	err = json.Unmarshal(b, &consumerKeys)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToUnmarshalConsumer,
			Error:   err.Error(),
		}
	}

	return consumerKeys, nil
}

func (a Auth) RetrieveConsumerKey(consumerId string) (*model.ConsumerKey, *errors.ApiError) {
	var consumerKey *model.ConsumerKey
	keys, apiError := a.GetAllConsumerKeys(consumerId)
	if apiError != nil {
		return nil, apiError
	}

	if keys == nil || len(keys.Data) == 0 {
		consumerKey, apiError = a.CreateConsumerKey(consumerId)
		if apiError != nil {
			return nil, apiError
		}
		return consumerKey, nil
	}

	return &keys.Data[0], nil
}

func (a Auth) DeleteConsumer(concatUrl string) *errors.ApiError {
	err := a.HttpRequest.DeleteRequest(concatUrl)
	if err != nil {
		logger.LogError(err)
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: authErrors.FailedToDeleteConsumer,
			Error:   err.Error(),
		}
	}
	return nil
}
