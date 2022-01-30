package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

// RegistryModule process the request recieved from handler
type RegistryModule struct {
	Repository registries.Repository
	Log        logger.LogInterface
}

func NewRegistryModule(repo registries.Repository, log logger.LogInterface) RegistryModule {
	return RegistryModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all registries.
func (m RegistryModule) Get() ([]registries.Registry, *errors.ApiError) {
	return m.Repository.Get()
}

// GetByFilter returns registries filtered
func (m RegistryModule) GetByFilter(filter registries.Filter) ([]registries.Registry, *errors.ApiError) {
	return m.Repository.GetByFilter(filter)
}

// Find return one registry by ID.
func (m RegistryModule) Find(id string) (registries.Registry, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a registry to the database.
func (m RegistryModule) Create(registry registries.Registry) (*registries.Registry, *errors.ApiError) {
	return m.Repository.Create(registry)
}

// Update update an existent registry.
func (m RegistryModule) Update(id string, upRegistry registries.Registry) *errors.ApiError {
	return m.Repository.Update(id, upRegistry)
}

// Delete delete an existent registry.
func (m RegistryModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
