package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
)

type LendingModuleInterface interface {
	Get() ([]model.Lending, *errors.ApiError)
	Find(id string) (model.Lending, *errors.ApiError)
	Create(lending model.Lending) (uint, *errors.ApiError)
	Update(id string, upLending model.Lending) *errors.ApiError
	Delete(id string) *errors.ApiError
}
