package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/model"
)

type ReserveRepositoryInterface interface {
	Get(bookId string) ([]model.Reserve, *errors.ApiError)
	Find(id string) (model.Reserve, *errors.ApiError)
	Create(reserve model.Reserve) (uint, *errors.ApiError)
	Update(id string, upReserve model.Reserve) *errors.ApiError
	Delete(id string) *errors.ApiError
}
