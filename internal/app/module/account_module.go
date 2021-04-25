package module

import (
	"fmt"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
	"github.com/FelipeAz/golibcontrol/platform/jwt"
	"github.com/FelipeAz/golibcontrol/platform/redis"
)

type AccountModule struct {
	Repository repository.AccountRepository
	Cache      *redis.Cache
}

// Login authenticate the user
func (m AccountModule) Login(credentials model.Account) (message login.Message) {
	account, apiError := m.Repository.Login(credentials)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Reason:  apiError.Error,
		}
	}

	token, apiError := jwt.CreateToken(account.ID)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Reason:  apiError.Error,
		}
	}

	apiError = m.Cache.CreateAuth(account.ID, token)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Reason:  apiError.Error,
		}
	}

	message.Token = token.AccessToken
	message.Message = fmt.Sprintf(login.SuccessMessage, account.FirstName)

	return
}

// Get returns all accounts.
func (m AccountModule) Get() (accounts []model.Account, apiError *errors.ApiError) {
	accounts, apiError = m.Repository.Get()
	return
}

// Find return one account by ID.
func (m AccountModule) Find(id string) (account model.Account, apiError *errors.ApiError) {
	account, apiError = m.Repository.Find(id)
	return
}

// Create creates an account
func (m AccountModule) Create(account model.Account) (id uint, apiError *errors.ApiError) {
	id, apiError = m.Repository.Create(account)
	return
}

// Update update an existent account.
func (m AccountModule) Update(id string, upAccount model.Account) (account model.Account, apiError *errors.ApiError) {
	account, apiError = m.Repository.Update(id, upAccount)
	return
}

// Delete delete an existent account by id.
func (m AccountModule) Delete(id string) (apiError *errors.ApiError) {
	apiError = m.Repository.Delete(id)
	return
}
