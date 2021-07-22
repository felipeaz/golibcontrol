package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model/converter"
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
	result, apiError := r.DB.Get(&[]model.Account{})
	if apiError != nil {
		return nil, apiError
	}
	accounts, apiError := converter.ConvertToSliceAccountObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return accounts, nil
}

// Find return one user by ID.
func (r AccountRepository) Find(id string) (model.Account, *errors.ApiError) {
	result, apiError := r.DB.Find(&model.Account{}, id)
	if apiError != nil {
		return model.Account{}, apiError
	}
	account, apiError := converter.ConvertToAccountObj(result)
	if apiError != nil {
		return model.Account{}, apiError
	}
	return account, nil
}

// FindWhere user by field and value.
func (r AccountRepository) FindWhere(fieldName, fieldValue string) (model.Account, *errors.ApiError) {
	result, apiError := r.DB.FindWhere(&model.Account{}, fieldName, fieldValue)
	if apiError != nil {
		return model.Account{}, apiError
	}
	account, apiError := converter.ConvertToAccountObj(result)
	if apiError != nil {
		return model.Account{}, apiError
	}
	return account, nil
}

// Create creates an user
func (r AccountRepository) Create(account model.Account) (uint, *errors.ApiError) {
	apiError := r.DB.Create(&account)
	if apiError != nil {
		return 0, apiError
	}
	return account.ID, nil
}

// Update update an existent user.
func (r AccountRepository) Update(id string, upAccount model.Account) *errors.ApiError {
	return r.DB.Update(&upAccount, id)
}

// Delete delete an existent user by id.
func (r AccountRepository) Delete(id string) *errors.ApiError {
	return r.DB.Delete(&model.Account{}, id)
}
