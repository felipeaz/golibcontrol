package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/category/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/category/model/converter"
)

// CategoryRepository is responsible of getting/saving information from DB.
type CategoryRepository struct {
	DB database.GORMServiceInterface
}

// Get returns all categories.
func (r CategoryRepository) Get() ([]model.Category, *errors.ApiError) {
	result, apiError := r.DB.Get(&model.Category{})
	if apiError != nil {
		return nil, apiError
	}
	categories, apiError := converter.ConvertToSliceCategoryObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return categories, nil
}

// Find return one category from DB by ID.
func (r CategoryRepository) Find(id string) (model.Category, *errors.ApiError) {
	result, apiError := r.DB.Find(model.Category{}, id)
	if apiError != nil {
		return model.Category{}, nil
	}
	categoryObj, apiError := converter.ConvertToCategoryObj(result)
	if apiError != nil {
		return model.Category{}, apiError
	}
	return categoryObj, nil
}

// Create persist a category to the DB.
func (r CategoryRepository) Create(category model.Category) (uint, *errors.ApiError) {
	apiError := r.DB.Create(&category)
	if apiError != nil {
		return 0, apiError
	}
	return category.ID, nil
}

// Update update an existent category.
func (r CategoryRepository) Update(id string, upCategory model.Category) (model.Category, *errors.ApiError) {
	result, apiError := r.DB.Update(upCategory, id)
	if apiError != nil {
		return model.Category{}, apiError
	}
	categoryObj, apiError := converter.ConvertToCategoryObj(result)
	if apiError != nil {
		return model.Category{}, apiError
	}
	return categoryObj, nil
}

// Delete delete an existent category from DB.
func (r CategoryRepository) Delete(id string) *errors.ApiError {
	apiError := r.DB.Delete(&model.Category{}, id)
	return apiError
}
