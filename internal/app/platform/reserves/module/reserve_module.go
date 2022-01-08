package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type ReserveModule struct {
	Repository reserves.Repository
	Log        logger.LogInterface
}

func NewReserveModule(repo reserves.Repository, log logger.LogInterface) ReserveModule {
	return ReserveModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ReserveModule) Get() ([]reserves.Reserve, *errors.ApiError) {
	return m.Repository.Get()
}

func (m ReserveModule) Find(id string) (reserves.Reserve, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ReserveModule) Create(comment reserves.Reserve) (*reserves.Reserve, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReserveModule) Update(id string, upReserve reserves.Reserve) *errors.ApiError {
	return m.Repository.Update(id, upReserve)
}

func (m ReserveModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
