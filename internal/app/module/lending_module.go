package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/interfaces/repository"
)

// LendingModule process the request recieved from handler.
type LendingModule struct {
	Repository repository.LendingRepositoryInterface
}

// Get returns all lendings.
func (m LendingModule) Get() (lendings []model.Lending, apiError *errors.ApiError) {
	lendings, apiError = m.Repository.Get()
	return
}

// Find return one lending by ID.
func (m LendingModule) Find(id string) (lending model.Lending, apiError *errors.ApiError) {
	lending, apiError = m.Repository.Find(id)
	return
}

// Create persist a lending to the database.
func (m LendingModule) Create(lending model.Lending) (id uint, apiError *errors.ApiError) {
	id, apiError = m.Repository.Create(lending)
	return
}

// Update update an existent lending.
func (m LendingModule) Update(id string, upLending model.Lending) (lending model.Lending, apiError *errors.ApiError) {
	lending, apiError = m.Repository.Update(id, upLending)
	return
}

// Delete delete an existent lending.
func (m LendingModule) Delete(id string) (apiError *errors.ApiError) {
	apiError = m.Repository.Delete(id)
	return
}
