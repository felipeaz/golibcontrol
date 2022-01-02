package module

import (
	"encoding/json"
	"fmt"
	"github.com/FelipeAz/golibcontrol/infra/consumer/jwt"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/consumer"
	databaseInterface "github.com/FelipeAz/golibcontrol/internal/app/database"
	domain "github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	user "github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
	"net/http"
)

type AuthModule struct {
	Repository user.Repository
	Consumer   consumer.Interface
	Cache      databaseInterface.CacheInterface
	Log        logger.LogInterface
}

func NewAuthModule(
	repo user.Repository,
	consumer consumer.Interface,
	cache databaseInterface.CacheInterface,
	log logger.LogInterface,
) AuthModule {
	return AuthModule{
		Repository: repo,
		Consumer:   consumer,
		Cache:      cache,
		Log:        log,
	}
}

// Login authenticate the user
func (m AuthModule) Login(credentials users.Account) login.Message {
	userName, token, consumerId, userId, apiError := m.authUser(credentials)
	if apiError != nil {
		return login.Message{
			Status:     apiError.Status,
			Message:    apiError.Message,
			Reason:     apiError.Error,
			UserId:     userId,
			ConsumerId: consumerId,
		}
	}

	return login.Message{
		Message: fmt.Sprintf(login.SuccessMessage, userName),
		Token:   token,
	}
}

// Logout authenticate the user
func (m AuthModule) Logout(session domain.Session) (message login.Message) {
	data, err := m.Cache.Get(session.ConsumerId)
	if err != nil {
		return login.Message{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToGetAuthenticationOnCache,
			Reason:  err.Error(),
		}
	}

	var consumerKeyId string
	switch data {
	case nil:
		consumerKey, err := m.Consumer.RetrieveConsumerKey(session.ConsumerId)
		if err != nil {
			return login.Message{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		consumerKeyId = consumerKey.Id
	default:
		var userAuth domain.Session
		err = json.Unmarshal(data, &userAuth)
		if err != nil {
			return login.Message{
				Status:  http.StatusInternalServerError,
				Message: errors.FailedToParseAuthenticationFromCache,
				Reason:  err.Error(),
			}
		}
		consumerKeyId = userAuth.ConsumerKeyId
	}

	err = m.Consumer.DeleteConsumerKey(session.ConsumerId, consumerKeyId)
	if err != nil {
		return login.Message{
			Status: http.StatusInternalServerError,
			Reason: err.Error(),
		}
	}

	err = m.Cache.Flush(session.ConsumerId)
	if err != nil {
		return login.Message{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToDeleteAuthenticationOnCache,
			Reason:  err.Error(),
		}
	}

	return login.Message{
		Status:  http.StatusOK,
		Message: login.LogoutSuccessMessage,
	}
}

// authUser retrieves user and authorize the access if the credentials match
func (m AuthModule) authUser(credentials users.Account) (string, string, string, uint, *errors.ApiError) {
	account, apiError := m.Repository.FindWhere("email", credentials.Email)
	if apiError != nil {
		return "", "", "", 0, &errors.ApiError{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Error:   login.AccountNotFoundMessage,
		}
	}

	if account.Password != credentials.Password {
		return "", "", "", 0, &errors.ApiError{
			Status:  http.StatusUnauthorized,
			Message: login.FailMessage,
			Error:   login.InvalidPasswordMessage,
		}
	}

	consumerKey, err := m.Consumer.RetrieveConsumerKey(account.ConsumerId)
	if err != nil {
		return "", "", "", 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToRetrieveConsumerKey,
			Error:   err.Error(),
		}
	}

	token, err := jwt.CreateToken(account.Email, consumerKey.Key, consumerKey.Secret)
	if err != nil {
		return "", "", "", 0, &errors.ApiError{
			Error: err.Error(),
		}
	}

	data := domain.Session{
		ConsumerId:    account.ConsumerId,
		ConsumerKeyId: consumerKey.Id,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return "", "", "", 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToMarshalAuthenticationOnCache,
			Error:   err.Error(),
		}
	}

	err = m.Cache.Set(account.ConsumerId, b)
	if err != nil {
		return "", "", "", 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToStoreAuthenticationKeyOnCache,
			Error:   err.Error(),
		}
	}

	return account.FirstName, token, account.ConsumerId, account.ID, nil
}
