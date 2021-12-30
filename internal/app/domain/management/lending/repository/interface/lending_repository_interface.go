package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
)

type LendingRepositoryInterface interface {
	Get() (lendings []model.Lending, apiError *errors.ApiError)
	Find(id string) (lending model.Lending, apiError *errors.ApiError)
	Create(lending model.Lending) (*model.Lending, *errors.ApiError)
	Update(id string, upLending model.Lending) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
