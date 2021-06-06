package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
)

type AccountRepositoryInterface interface {
	Get() (accounts []model.Account, apiError *errors.ApiError)
	Find(id string) (account model.Account, apiError *errors.ApiError)
	FindWhere(fieldName, fieldValue string) (account model.Account, apiError *errors.ApiError)
	Create(account model.Account) (uint, *errors.ApiError)
	Update(id string, upAccount model.Account) (model.Account, *errors.ApiError)
	Delete(id string) (apiError *errors.ApiError)
}
