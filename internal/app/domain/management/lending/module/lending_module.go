package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// LendingModule process the request received from handler.
type LendingModule struct {
	Repository _interface.LendingRepositoryInterface
	Log        logger.LogInterface
}

func NewLendingModule(
	repo _interface.LendingRepositoryInterface,
	log logger.LogInterface) LendingModule {
	return LendingModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all lending.
func (m LendingModule) Get() ([]model.Lending, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one lending by ID.
func (m LendingModule) Find(id string) (model.Lending, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a lending to the database.
func (m LendingModule) Create(lending model.Lending) (*model.Lending, *errors.ApiError) {
	return m.Repository.Create(lending)
}

// Update update an existent lending.
func (m LendingModule) Update(id string, upLending model.Lending) *errors.ApiError {
	return m.Repository.Update(id, upLending)
}

// Delete delete an existent lending.
func (m LendingModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
