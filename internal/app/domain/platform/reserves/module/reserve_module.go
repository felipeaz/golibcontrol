package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type ReserveModule struct {
	Repository _interface.ReserveRepositoryInterface
	Log        logger.LogInterface
}

func NewReserveModule(repo _interface.ReserveRepositoryInterface, log logger.LogInterface) ReserveModule {
	return ReserveModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ReserveModule) Get() ([]model.Reserve, *errors.ApiError) {
	return m.Repository.Get()
}

func (m ReserveModule) Find(id string) (model.Reserve, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ReserveModule) Create(comment model.Reserve) (*model.Reserve, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReserveModule) Update(id string, upReserve model.Reserve) *errors.ApiError {
	return m.Repository.Update(id, upReserve)
}

func (m ReserveModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
