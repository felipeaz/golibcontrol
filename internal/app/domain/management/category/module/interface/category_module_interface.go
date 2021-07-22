package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/model"
)

type CategoryModuleInterface interface {
	Get() ([]model.Category, *errors.ApiError)
	Find(id string) (model.Category, *errors.ApiError)
	Create(category model.Category) (uint, *errors.ApiError)
	Update(id string, upCategory model.Category) *errors.ApiError
	Delete(id string) *errors.ApiError
}
