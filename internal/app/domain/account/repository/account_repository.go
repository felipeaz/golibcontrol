package repository

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/platform/logger"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

// Get returns all accounts.
func (r AccountRepository) Get() (accounts []model.Account, apiError *errors.ApiError) {
	result := r.DB.Find(&accounts)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	return
}

// Find return one account by ID.
func (r AccountRepository) Find(id string) (account model.Account, apiError *errors.ApiError) {
	result := r.DB.Model(model.Account{}).Where("id = ?", id).First(&account)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return model.Account{}, &errors.ApiError{
				Status:  http.StatusNotFound,
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}

		return model.Account{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "account not found",
		}
	}
	return
}

// Create creates an account
func (r AccountRepository) Create(account model.Account) (id uint, apiError *errors.ApiError) {
	result := r.DB.Create(&account)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return account.ID, nil
}

// Update update an existent account.
func (r AccountRepository) Update(id string, upAccount model.Account) (account model.Account, apiError *errors.ApiError) {
	account, apiError = r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return
	}

	result := r.DB.Model(&account).Updates(&upAccount)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return model.Account{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}

	return
}

// Delete delete an existent account by id.
func (r AccountRepository) Delete(id string) (apiError *errors.ApiError) {
	account, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailMessage
		return
	}

	err := r.DB.Delete(&account).Error
	if err != nil {
		logger.LogError(err)
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}

	return
}

// Login authenticate user if credentials are right
func (r AccountRepository) Login(credentials model.Account) (account model.Account, apiError *errors.ApiError) {
	result := r.DB.Model(model.Account{}).
		Where("email = ?", credentials.Email).First(&account)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return model.Account{}, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: login.FailMessage,
				Error:   err.Error(),
			}
		}

		return model.Account{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: login.FailMessage,
			Error:   login.AccountNotFoundMessage,
		}
	}

	// Validate Password
	if account.Password != credentials.Password {
		return model.Account{}, &errors.ApiError{
			Status:  http.StatusUnauthorized,
			Message: login.FailMessage,
			Error:   login.InvalidPasswordMessage,
		}
	}

	return
}
