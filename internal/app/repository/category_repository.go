package repository

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/platform/logger"
)

// CategoryRepository is responsible of getting/saving information from DB.
type CategoryRepository struct {
	DB *gorm.DB
}

// Get returns all categories.
func (r CategoryRepository) Get() (categories []model.Category, apiError *errors.ApiError) {
	result := r.DB.Find(&categories)
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

// Find return one category from DB by ID.
func (r CategoryRepository) Find(id string) (category model.Category, apiError *errors.ApiError) {
	result := r.DB.Model(model.Category{}).Where("id = ?", id).First(&category)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return model.Category{}, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}

		return model.Category{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "category not found",
		}
	}

	return
}

// Create persist a category to the DB.
func (r CategoryRepository) Create(category model.Category) (uint, *errors.ApiError) {
	result := r.DB.Create(&category)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return category.ID, nil
}

// Update update an existent category.
func (r CategoryRepository) Update(id string, upCategory model.Category) (model.Category, *errors.ApiError) {
	category, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return model.Category{}, apiError
	}

	result := r.DB.Model(&category).Updates(upCategory)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return model.Category{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}

	return category, nil
}

// Delete delete an existent category from DB.
func (r CategoryRepository) Delete(id string) (apiError *errors.ApiError) {
	category, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailMessage
		return
	}

	err := r.DB.Delete(&category).Error
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
