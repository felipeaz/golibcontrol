package module

import (
	"encoding/json"
	"fmt"
	"github.com/FelipeAz/golibcontrol/infra/consumer/jwt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/consumer"
	databaseInterface "github.com/FelipeAz/golibcontrol/internal/database"
	"github.com/FelipeAz/golibcontrol/internal/logger"
	"net/http"
	"strings"
)

type AuthModule struct {
	Repository users.Repository
	Consumer   consumer.Interface
	Cache      databaseInterface.Cache
	Log        logger.LogInterface
}

func NewAuthModule(
	repo users.Repository,
	consumer consumer.Interface,
	cache databaseInterface.Cache,
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
func (m AuthModule) Login(credentials users.Account) login.Data {
	userName, token, apiError := m.authUser(credentials)
	if apiError != nil {
		return login.Data{
			Status:  apiError.Status,
			Message: apiError.Message,
			Reason:  apiError.Error,
		}
	}

	return login.Data{
		Message: fmt.Sprintf(login.SuccessMessage, strings.Title(userName)),
		Token:   token,
	}
}

// Logout authenticate the user
func (m AuthModule) Logout(session auth.Session) (message login.Data) {
	data, err := m.Cache.Get(session.ConsumerId)
	if err != nil {
		return login.Data{
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
			return login.Data{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		consumerKeyId = consumerKey.Id
	default:
		var userAuth auth.Session
		err = json.Unmarshal(data, &userAuth)
		if err != nil {
			return login.Data{
				Status:  http.StatusInternalServerError,
				Message: errors.FailedToParseAuthenticationFromCache,
				Reason:  err.Error(),
			}
		}
		consumerKeyId = userAuth.ConsumerKeyId
	}

	err = m.Consumer.DeleteConsumerKey(session.ConsumerId, consumerKeyId)
	if err != nil {
		return login.Data{
			Status: http.StatusInternalServerError,
			Reason: err.Error(),
		}
	}

	err = m.Cache.Flush(session.ConsumerId)
	if err != nil {
		return login.Data{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToDeleteAuthenticationOnCache,
			Reason:  err.Error(),
		}
	}

	return login.Data{
		Status:  http.StatusOK,
		Message: login.LogoutSuccessMessage,
	}
}

// authUser retrieves user and authorize the access if the credentials match
func (m AuthModule) authUser(credentials users.Account) (string, string, *errors.ApiError) {
	account, apiError := m.Repository.FindWhere("email", credentials.Email)
	if apiError != nil {
		return "", "", &errors.ApiError{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Error:   login.AccountNotFoundMessage,
		}
	}

	if account.Password != credentials.Password {
		return "", "", &errors.ApiError{
			Status:  http.StatusUnauthorized,
			Message: login.FailMessage,
			Error:   login.InvalidPasswordMessage,
		}
	}

	consumerKey, err := m.Consumer.RetrieveConsumerKey(account.ConsumerId)
	if err != nil {
		return "", "", &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToRetrieveConsumerKey,
			Error:   err.Error(),
		}
	}

	token, err := jwt.CreateToken(account.ID, account.Email, consumerKey.Key, consumerKey.Secret)
	if err != nil {
		return "", "", &errors.ApiError{
			Error: err.Error(),
		}
	}

	data := auth.Session{
		ConsumerId:    account.ConsumerId,
		ConsumerKeyId: consumerKey.Id,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return "", "", &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToMarshal,
			Error:   err.Error(),
		}
	}

	err = m.Cache.Set(account.ConsumerId, b)
	if err != nil {
		return "", "", &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToStoreAuthenticationKeyOnCache,
			Error:   err.Error(),
		}
	}

	return account.FirstName, token, nil
}
