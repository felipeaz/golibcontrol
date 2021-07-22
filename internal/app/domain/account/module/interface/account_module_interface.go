package _interface

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
)

type AccountModuleInterface interface {
	Login(credentials model.Account) login.Message
	StoreAuthUser(account model.Account) login.Message
	Logout(r *http.Request) (message login.Message)
	Get() ([]model.Account, *errors.ApiError)
	Find(id string) (model.Account, *errors.ApiError)
	Create(account model.Account) (uint, *errors.ApiError)
	Update(id string, upAccount model.Account) *errors.ApiError
	Delete(id string) *errors.ApiError
}
