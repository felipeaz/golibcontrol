package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	domain "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// LendingModule process the request received from handler.
type LendingModule struct {
	Repository domain.Repository
	Log        logger.LogInterface
}

func NewLendingModule(
	repo domain.Repository,
	log logger.LogInterface) LendingModule {
	return LendingModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all lending.
func (m LendingModule) Get() ([]domain.Lending, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one lending by ID.
func (m LendingModule) Find(id string) (domain.Lending, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a lending to the database.
func (m LendingModule) Create(lending domain.Lending) (*domain.Lending, *errors.ApiError) {
	return m.Repository.Create(lending)
}

// Update update an existent lending.
func (m LendingModule) Update(id string, upLending domain.Lending) *errors.ApiError {
	return m.Repository.Update(id, upLending)
}

// Delete delete an existent lending.
func (m LendingModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
