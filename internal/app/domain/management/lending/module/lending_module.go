package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/repository/interface"
)

// LendingModule process the request recieved from handler.
type LendingModule struct {
	Repository _interface.LendingRepositoryInterface
}

func NewLendingModule(repo _interface.LendingRepositoryInterface) LendingModule {
	return LendingModule{
		Repository: repo,
	}
}

// Get returns all lendings.
func (m LendingModule) Get() ([]model.Lending, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one lending by ID.
func (m LendingModule) Find(id string) (model.Lending, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a lending to the database.
func (m LendingModule) Create(lending model.Lending) (uint, *errors.ApiError) {
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
