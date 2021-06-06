package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/category/model"
)

type CategoryRepositoryInterface interface {
	Get() (categories []model.Category, apiError *errors.ApiError)
	Find(id string) (category model.Category, apiError *errors.ApiError)
	Create(category model.Category) (uint, *errors.ApiError)
	Update(id string, upCategory model.Category) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
