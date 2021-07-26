package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/repository/interface"
)

type ReserveModule struct {
	Repository _interface.ReserveRepositoryInterface
}

func NewReserveModule(repo _interface.ReserveRepositoryInterface) ReserveModule {
	return ReserveModule{
		Repository: repo,
	}
}

func (m ReserveModule) Get(bookId string) ([]model.Reserve, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m ReserveModule) Find(id string) (model.Reserve, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ReserveModule) Create(comment model.Reserve) (uint, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReserveModule) Update(id string, upReserve model.Reserve) *errors.ApiError {
	return m.Repository.Update(id, upReserve)
}

func (m ReserveModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
