package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/model"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

// Get returns all accounts.
func (r *AccountRepositoryMock) Get() ([]model.Account, *errors.ApiError) {
	resp := r.Called()
	return resp.Get(0).([]model.Account), resp.Get(1).(*errors.ApiError)
}

// Find return one user by ID.
func (r *AccountRepositoryMock) Find(id string) (model.Account, *errors.ApiError) {
	resp := r.Called(id)
	return resp.Get(0).(model.Account), resp.Get(1).(*errors.ApiError)
}

// FindWhere user by field and value.
func (r *AccountRepositoryMock) FindWhere(fieldName, fieldValue string) (model.Account, *errors.ApiError) {
	resp := r.Called(fieldName, fieldValue)
	return resp.Get(0).(model.Account), resp.Get(1).(*errors.ApiError)
}

// Create creates a user
func (r *AccountRepositoryMock) Create(account model.Account) (uint, *errors.ApiError) {
	resp := r.Called(account)
	return resp.Get(0).(uint), resp.Get(1).(*errors.ApiError)
}

// Update update an existent user.
func (r *AccountRepositoryMock) Update(id string, upAccount model.Account) *errors.ApiError {
	resp := r.Called(id, upAccount)
	return resp.Get(0).(*errors.ApiError)
}

// Delete delete an existent user by id.
func (r *AccountRepositoryMock) Delete(id string) *errors.ApiError {
	resp := r.Called(id)
	return resp.Get(0).(*errors.ApiError)
}
