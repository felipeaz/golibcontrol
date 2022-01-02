package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/pkg"
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
func (r AccountRepository) Get() ([]users.Account, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]users.Account{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	accounts, apiError := pkg.ParseInterfaceToSliceAccount(result)
	if apiError != nil {
		return nil, apiError
	}
	return accounts, nil
}

// Find return one user by ID.
func (r AccountRepository) Find(id string) (users.Account, *errors.ApiError) {
	result, err := r.DB.Fetch(&users.Account{}, id)
	if err != nil {
		return users.Account{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	account, apiError := pkg.ParseInterfaceToAccount(result)
	if apiError != nil {
		return users.Account{}, apiError
	}
	return account, nil
}

// FindWhere user by field and value.
func (r AccountRepository) FindWhere(fieldName, fieldValue string) (users.Account, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&users.Account{}, fieldName, fieldValue)
	if err != nil {
		return users.Account{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	account, apiError := pkg.ParseInterfaceToAccount(result)
	if apiError != nil {
		return users.Account{}, apiError
	}
	return account, nil
}

// Create creates a user
func (r AccountRepository) Create(account users.Account) (*users.Account, *errors.ApiError) {
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
func (r AccountRepository) Update(id string, upAccount users.Account) *errors.ApiError {
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
	err := r.DB.Remove(&users.Account{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
