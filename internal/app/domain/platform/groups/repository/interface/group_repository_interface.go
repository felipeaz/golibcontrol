package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/model"
)

type GroupRepositoryInterface interface {
	Get() ([]model.Group, *errors.ApiError)
	Find(id string) (model.Group, *errors.ApiError)
	Create(group model.Group) (uint, *errors.ApiError)
	Update(id string, upGroup model.Group) *errors.ApiError
	Delete(id string) *errors.ApiError
}
