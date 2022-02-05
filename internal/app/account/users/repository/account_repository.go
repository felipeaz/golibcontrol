package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
	result, err := r.DB.Find(nil, &[]users.Account{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseInterfaceToSliceAccount(result)
}

// Find return one user by ID.
func (r AccountRepository) Find(id string) (users.Account, *errors.ApiError) {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &users.Account{})
	if err != nil {
		return users.Account{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseInterfaceToAccount(result)
}

// FindWhere user by field and value.
func (r AccountRepository) FindWhere(fieldName, fieldValue string) (users.Account, *errors.ApiError) {
	tx := r.DB.Where(nil, fmt.Sprintf("`%s`='%s'", fieldName, fieldValue))
	result, err := r.DB.Find(tx, &users.Account{})
	if err != nil {
		return users.Account{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseInterfaceToAccount(result)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upAccount)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	err = r.DB.Set(tx, &upAccount, "student_account", upAccount.StudentAccount)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &users.Account{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
