package module

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/infra/auth/jwt"
	"github.com/FelipeAz/golibcontrol/internal/app/auth"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	databaseInterface "github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/repository/interface"
)

type AccountModule struct {
	Repository _interface.AccountRepositoryInterface
	Auth       auth.AuthInterface
	Cache      databaseInterface.CacheInterface
}

func NewAccountModule(repo _interface.AccountRepositoryInterface, auth auth.AuthInterface, cache databaseInterface.CacheInterface) AccountModule {
	return AccountModule{
		Repository: repo,
		Auth:       auth,
		Cache:      cache,
	}
}

// Login authenticate the user
func (m AccountModule) Login(credentials model.Account) login.Message {
	userName, token, apiError := m.authUser(credentials)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: apiError.Message,
			Reason:  apiError.Error,
		}
	}

	return login.Message{
		Message: fmt.Sprintf(login.SuccessMessage, userName),
		Token:   token,
	}
}

// Logout authenticate the user
func (m AccountModule) Logout(session model.UserSession) (message login.Message) {
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
		consumerKey, err := m.Auth.RetrieveConsumerKey(session.ConsumerId, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			return login.Message{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		consumerKeyId = consumerKey.Id
	default:
		var userAuth model.UserSession
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

	err = m.Auth.DeleteConsumer(session.ConsumerId, consumerKeyId)
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

// Get returns all accounts.
func (m AccountModule) Get() ([]model.Account, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one user by ID.
func (m AccountModule) Find(id string) (model.Account, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create creates a user
func (m AccountModule) Create(account model.Account) (uint, *errors.ApiError) {
	consumer, err := m.Auth.CreateConsumer(account.Email)
	if err != nil {
		return 0, &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToCreateConsumer,
			Error:   err.Error(),
		}
	}
	account.ConsumerId = consumer.Id

	return m.Repository.Create(account)
}

// Update update an existent user.
func (m AccountModule) Update(id string, upAccount model.Account) *errors.ApiError {
	return m.Repository.Update(id, upAccount)
}

// Delete delete an existent user by id.
func (m AccountModule) Delete(id string) *errors.ApiError {
	user, apiError := m.Repository.Find(id)
	if apiError != nil {
		return apiError
	}

	err := m.Auth.DeleteConsumer(user.ConsumerId, "")
	if err != nil {
		return &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToDeleteConsumer,
			Error:   err.Error(),
		}
	}

	err = m.Cache.Flush(user.ConsumerId)
	if err != nil {
		return &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToDeleteAuthenticationOnCache,
			Error:   err.Error(),
		}
	}

	return m.Repository.Delete(id)
}

// authUser retrieves user and authorize the access if the credentials match
func (m AccountModule) authUser(credentials model.Account) (string, string, *errors.ApiError) {
	account, apiError := m.Repository.FindWhere("email", credentials.Email)
	if apiError != nil {
		return "", "", &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  apiError.Status,
			Message: login.FailMessage,
			Error:   login.AccountNotFoundMessage,
		}
	}

	if account.Password != credentials.Password {
		return "", "", &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusUnauthorized,
			Message: login.FailMessage,
			Error:   login.InvalidPasswordMessage,
		}
	}

	consumerKey, err := m.Auth.RetrieveConsumerKey(account.ConsumerId, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return "", "", &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToRetrieveConsumerKey,
			Error:   err.Error(),
		}
	}

	token, err := jwt.CreateToken(account.Email, consumerKey.Key, consumerKey.Secret)
	if err != nil {
		return "", "", &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Error:   err.Error(),
		}
	}

	data := model.UserSession{
		ConsumerId:    account.ConsumerId,
		ConsumerKeyId: consumerKey.Id,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return "", "", &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToMarshalAuthenticationOnCache,
			Error:   err.Error(),
		}
	}

	err = m.Cache.Set(account.ConsumerId, b)
	if err != nil {
		return "", "", &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToStoreAuthenticationKeyOnCache,
			Error:   err.Error(),
		}
	}

	return account.FirstName, token, nil
}
