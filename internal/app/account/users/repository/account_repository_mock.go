package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

// Get returns all accounts.
func (r *AccountRepositoryMock) Get() ([]users.Account, *errors.ApiError) {
	resp := r.Called()
	return resp.Get(0).([]users.Account), resp.Get(1).(*errors.ApiError)
}

// Find return one user by ID.
func (r *AccountRepositoryMock) Find(id string) (users.Account, *errors.ApiError) {
	resp := r.Called(id)
	return resp.Get(0).(users.Account), resp.Get(1).(*errors.ApiError)
}

// FindWhere user by field and value.
func (r *AccountRepositoryMock) FindWhere(fieldName, fieldValue string) (users.Account, *errors.ApiError) {
	resp := r.Called(fieldName, fieldValue)
	return resp.Get(0).(users.Account), resp.Get(1).(*errors.ApiError)
}

// Create creates a user
func (r *AccountRepositoryMock) Create(account users.Account) (*users.Account, *errors.ApiError) {
	resp := r.Called(account)
	return resp.Get(0).(*users.Account), resp.Get(1).(*errors.ApiError)
}

// Update update an existent user.
func (r *AccountRepositoryMock) Update(id string, upAccount users.Account) *errors.ApiError {
	resp := r.Called(id, upAccount)
	return resp.Get(0).(*errors.ApiError)
}

// Delete delete an existent user by id.
func (r *AccountRepositoryMock) Delete(id string) *errors.ApiError {
	resp := r.Called(id)
	return resp.Get(0).(*errors.ApiError)
}
