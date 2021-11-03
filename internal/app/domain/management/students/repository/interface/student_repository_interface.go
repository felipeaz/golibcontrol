package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students/model"
)

type StudentRepositoryInterface interface {
	Get() (students []model.Student, apiError *errors.ApiError)
	Find(id string) (student model.Student, apiError *errors.ApiError)
	Create(student model.Student) (string, *errors.ApiError)
	Update(id string, upStudent model.Student) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
