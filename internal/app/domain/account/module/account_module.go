package module

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/account/repository/interface"
	"github.com/FelipeAz/golibcontrol/platform/jwt"
	"github.com/FelipeAz/golibcontrol/platform/redis"
)

type AccountModule struct {
	Repository _interface.AccountRepositoryInterface
	Auth       *jwt.Auth
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

	token, apiError := m.Auth.CreateToken(account.ID)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Reason:  apiError.Error,
		}
	}

	apiError = m.Cache.StoreAuth(account.ID, token)
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

// Logout authenticate the user
func (m AccountModule) Logout(r *http.Request) (message login.Message) {
	userAuth, err := m.Auth.GetAuthUser(r)
	if err != nil {
		return login.Message{
			Status:  http.StatusBadRequest,
			Message: login.LogoutFailMessage,
			Reason:  err.Error(),
		}
	}

	err = m.Auth.FetchAuth(&userAuth)
	if err != nil {
		return login.Message{
			Status:  http.StatusBadRequest,
			Message: login.UserNotLoggedIn,
			Reason:  err.Error(),
		}
	}

	userId := strconv.FormatUint(uint64(userAuth.UserId), 10)
	err = m.Cache.Flush(userId)
	if err != nil {
		return login.Message{
			Status:  http.StatusUnauthorized,
			Message: login.LogoutFailMessage,
			Reason:  err.Error(),
		}
	}

	return login.Message{
		Status:  http.StatusOK,
		Message: login.LogoutSuccessMessage,
	}
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
