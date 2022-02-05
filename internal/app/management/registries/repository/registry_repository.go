package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/app/filters"
	"github.com/FelipeAz/golibcontrol/internal/app/management/registries/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
)

// RegistryRepository is responsible for getting/saving information from DB.
type RegistryRepository struct {
	DB database.GORMServiceInterface
}

func NewRegistryRepository(db database.GORMServiceInterface) RegistryRepository {
	return RegistryRepository{
		DB: db,
	}
}

// Get returns all registries.
func (r RegistryRepository) Get() ([]registries.Registry, *errors.ApiError) {
	result, err := r.DB.Find(nil, &[]registries.Registry{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceRegistryObj(result)
}

// GetByFilter filters by query string.
func (r RegistryRepository) GetByFilter(filter registries.Filter) ([]registries.Registry, *errors.ApiError) {
	queryString := filters.BuildQueryFromFilter(filter)

	tx := r.DB.Where(nil, queryString)
	result, err := r.DB.Find(tx, &[]registries.Registry{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceRegistryObj(result)
}

// Find return one registry from DB by ID.
func (r RegistryRepository) Find(id string) (registries.Registry, *errors.ApiError) {
	tx := r.DB.Where(nil, fmt.Sprintf("registry_number = %s", id))
	result, err := r.DB.FindOne(tx, &registries.Registry{})
	if err != nil {
		return registries.Registry{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToRegistryObj(result)
}

// Create persist a registry to the DB.
func (r RegistryRepository) Create(registry registries.Registry) (*registries.Registry, *errors.ApiError) {
	err := r.DB.Persist(&registry)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &registry, nil
}

// Update update an existent registry.
func (r RegistryRepository) Update(id string, upRegistry registries.Registry) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("registry_number = %s", id))
	err := r.DB.Refresh(tx, &upRegistry)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	err = r.DB.Set(tx, &upRegistry, "available", upRegistry.Available)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	err = r.DB.Set(tx, &upRegistry, "reserved", upRegistry.Reserved)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// Delete an existent registry from DB.
func (r RegistryRepository) Delete(id string) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("registry_number = %s", id))
	err := r.DB.Remove(tx, &registries.Registry{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
