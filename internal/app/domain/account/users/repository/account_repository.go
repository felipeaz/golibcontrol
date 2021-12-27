package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/model/converter"
)

type AccountRepository struct {
	DB database.GORMServiceInterface
}

func NewAccountRepository(dbService database.GORMServiceInterface) AccountRepository {
	return AccountRepository{
		DB: dbService,
	}
}

// Get returns all accounts.
func (r AccountRepository) Get() ([]model.Account, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]model.Account{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	accounts, apiError := converter.ConvertToSliceAccountObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return accounts, nil
}

// Find return one user by ID.
func (r AccountRepository) Find(id string) (model.Account, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Account{}, id)
	if err != nil {
		return model.Account{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	account, apiError := converter.ConvertToAccountObj(result)
	if apiError != nil {
		return model.Account{}, apiError
	}
	return account, nil
}

// FindWhere user by field and value.
func (r AccountRepository) FindWhere(fieldName, fieldValue string) (model.Account, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&model.Account{}, fieldName, fieldValue)
	if err != nil {
		return model.Account{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	account, apiError := converter.ConvertToAccountObj(result)
	if apiError != nil {
		return model.Account{}, apiError
	}
	return account, nil
}

// Create creates a user
func (r AccountRepository) Create(account model.Account) (*model.Account, *errors.ApiError) {
	err := r.DB.Persist(&account)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &account, nil
}

// Update update an existent user.
func (r AccountRepository) Update(id string, upAccount model.Account) *errors.ApiError {
	err := r.DB.Refresh(&upAccount, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// Delete delete an existent user by id.
func (r AccountRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Account{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
