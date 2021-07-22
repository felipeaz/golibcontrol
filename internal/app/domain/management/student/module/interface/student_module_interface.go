package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
)

type StudentModuleInterface interface {
	Get() ([]model.Student, *errors.ApiError)
	Find(id string) (model.Student, *errors.ApiError)
	Create(student model.Student) (string, *errors.ApiError)
	Update(id string, upStudent model.Student) *errors.ApiError
	Delete(id string) *errors.ApiError
}
