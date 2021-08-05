package module

import (
	"fmt"
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/auth"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/repository/interface"
)

type AccountModule struct {
	Repository _interface.AccountRepositoryInterface
	Auth       auth.Auth
}

func NewAccountModule(repo _interface.AccountRepositoryInterface, auth auth.Auth) AccountModule {
	return AccountModule{
		Repository: repo,
		Auth:       auth,
	}
}

// Login authenticate the user
func (m AccountModule) Login(credentials model.Account) login.Message {
	account, apiError := m.authUser(credentials)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: apiError.Message,
			Reason:  apiError.Error,
		}
	}

	consumerKey, apiError := m.Auth.RetrieveConsumerKey(account.ConsumerId)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: apiError.Message,
		}
	}

	return login.Message{
		Message: fmt.Sprintf(login.SuccessMessage, account.FirstName),
		Token:   consumerKey.Key,
	}
}

// Logout authenticate the user
func (m AccountModule) Logout(session model.UserSession) (message login.Message) {
	concatUrl := session.UserId + "/key-auth/" + session.KeyId
	apiError := m.Auth.DeleteConsumer(concatUrl)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: login.LogoutFailMessage,
			Reason:  apiError.Error,
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

// Create creates an user
func (m AccountModule) Create(account model.Account) (uint, *errors.ApiError) {
	consumer, apiError := m.Auth.CreateConsumer(account.Email)
	if apiError != nil {
		return 0, apiError
	}
	account.ConsumerId = consumer.Id

	consumerKey, apiError := m.Auth.CreateConsumerKey(account.ConsumerId)
	if apiError != nil {
		return 0, apiError
	}
	account.ConsumerKeyId = consumerKey.Id

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
	apiError = m.Auth.DeleteConsumer(user.ConsumerId)
	if apiError != nil {
		return apiError
	}
	return m.Repository.Delete(id)
}

// authUser retrieves user and authorize the access if the credentials match
func (m AccountModule) authUser(credentials model.Account) (model.Account, *errors.ApiError) {
	account, apiError := m.Repository.FindWhere("email", credentials.Email)
	if apiError != nil {
		return model.Account{}, &errors.ApiError{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Error:   login.AccountNotFoundMessage,
		}
	}

	if account.Password != credentials.Password {
		return model.Account{}, &errors.ApiError{
			Status:  http.StatusUnauthorized,
			Message: login.FailMessage,
			Error:   login.InvalidPasswordMessage,
		}
	}

	return account, nil
}
