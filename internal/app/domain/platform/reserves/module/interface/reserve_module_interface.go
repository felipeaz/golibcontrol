package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/model"
)

type ReserveModuleInterface interface {
	Get() ([]model.Reserve, *errors.ApiError)
	Find(id string) (model.Reserve, *errors.ApiError)
	Create(comment model.Reserve) (*model.Reserve, *errors.ApiError)
	Update(id string, upReserve model.Reserve) *errors.ApiError
	Delete(id string) *errors.ApiError
}
