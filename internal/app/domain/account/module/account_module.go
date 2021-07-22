package module

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/FelipeAz/golibcontrol/infra/jwt"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/repository/interface"
)

type AccountModule struct {
	Repository _interface.AccountRepositoryInterface
	Auth       *jwt.Auth
	Cache      *redis.Cache
}

func NewAccountModule(repo _interface.AccountRepositoryInterface, auth *jwt.Auth, cache *redis.Cache) AccountModule {
	return AccountModule{
		Repository: repo,
		Auth:       auth,
		Cache:      cache,
	}
}

// Login authenticate the user
func (m AccountModule) Login(credentials model.Account) login.Message {
	account, apiError := m.Repository.FindWhere("email", credentials.Email)
	if apiError != nil {
		return login.Message{
			Status:  apiError.Status,
			Message: login.FailMessage,
			Reason:  login.AccountNotFoundMessage,
		}
	}

	if account.Password != credentials.Password {
		return login.Message{
			Status:  http.StatusUnauthorized,
			Message: login.FailMessage,
			Reason:  login.InvalidPasswordMessage,
		}
	}

	message := m.StoreAuthUser(account)
	return message
}

// StoreAuthUser stores the authentication token on cache
func (m AccountModule) StoreAuthUser(account model.Account) login.Message {
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

	var message login.Message
	message.Token = token.AccessToken
	message.Message = fmt.Sprintf(login.SuccessMessage, account.FirstName)
	return message
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
func (m AccountModule) Get() ([]model.Account, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one user by ID.
func (m AccountModule) Find(id string) (model.Account, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create creates an user
func (m AccountModule) Create(account model.Account) (uint, *errors.ApiError) {
	return m.Repository.Create(account)
}

// Update update an existent user.
func (m AccountModule) Update(id string, upAccount model.Account) *errors.ApiError {
	return m.Repository.Update(id, upAccount)
}

// Delete delete an existent user by id.
func (m AccountModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
