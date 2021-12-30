package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/model"
)

type AccountModuleInterface interface {
	Login(credentials model.Account) login.Message
	Logout(session model.UserSession) (message login.Message)
	Get() ([]model.Account, *errors.ApiError)
	Find(id string) (model.Account, *errors.ApiError)
	Create(account model.Account) (*model.Account, *errors.ApiError)
	Update(id string, upAccount model.Account) *errors.ApiError
	Delete(id string) *errors.ApiError
}
