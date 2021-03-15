package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// LendingModule process the request recieved from handler.
type LendingModule struct {
	Repository repository.LendingRepository
}

// Get returns all lendings.
func (m LendingModule) Get() (lendings []model.Lending, err error) {
	lendings, err = m.Repository.Get()
	return
}

// Find return one lending by ID.
func (m LendingModule) Find(id int) (lending model.Lending, err error) {
	lending, err = m.Repository.Find(id)
	return
}

// Create persist a lending to the database.
func (m LendingModule) Create(lending model.Lending) (id uint, err error) {
	id, err = m.Repository.Create(lending)
	return
}

// Update update an existent lending.
func (m LendingModule) Update(id int, upLending model.Lending) (lending model.Lending, err error) {
	lending, err = m.Repository.Update(id, upLending)
	return
}

// Delete delete an existent lending.
func (m LendingModule) Delete(id int) (err error) {
	err = m.Repository.Delete(id)
	return
}
