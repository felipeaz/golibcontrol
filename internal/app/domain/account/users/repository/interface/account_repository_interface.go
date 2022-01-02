package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
)

type AccountRepositoryInterface interface {
	Get() (accounts []users.Account, apiError *errors.ApiError)
	Find(id string) (account users.Account, apiError *errors.ApiError)
	FindWhere(fieldName, fieldValue string) (account users.Account, apiError *errors.ApiError)
	Create(account users.Account) (*users.Account, *errors.ApiError)
	Update(id string, upAccount users.Account) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
