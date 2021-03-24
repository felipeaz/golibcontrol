package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type LendingRepositoryInterface interface {
	Get() (lendings []model.Lending, apiError *errors.ApiError)
	Find(id string) (lending model.Lending, apiError *errors.ApiError)
	Create(lending model.Lending) (uint, *errors.ApiError)
	Update(id string, upLending model.Lending) (model.Lending, *errors.ApiError)
	Delete(id string) (apiError *errors.ApiError)
	BeforeCreateAndUpdate(studentId, bookId uint) *errors.ApiError
	BeforeCreate(studentId, bookId uint) *errors.ApiError
}
