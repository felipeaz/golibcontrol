package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories/model/converter"
)

// CategoryRepository is responsible for getting/saving information from DB.
type CategoryRepository struct {
	DB database.GORMServiceInterface
}

func NewCategoryRepository(db database.GORMServiceInterface) CategoryRepository {
	return CategoryRepository{
		DB: db,
	}
}

// Get returns all categories.
func (r CategoryRepository) Get() ([]model.Category, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]model.Category{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	categories, apiError := converter.ConvertToSliceCategoryObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return categories, nil
}

// Find return one category from DB by ID.
func (r CategoryRepository) Find(id string) (model.Category, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Category{}, id)
	if err != nil {
		return model.Category{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	category, apiError := converter.ConvertToCategoryObj(result)
	if apiError != nil {
		return model.Category{}, apiError
	}

	return category, nil
}

// Create persist a category to the DB.
func (r CategoryRepository) Create(category model.Category) (uint, *errors.ApiError) {
	err := r.DB.Persist(&category)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return category.ID, nil
}

// Update update an existent category.
func (r CategoryRepository) Update(id string, upCategory model.Category) *errors.ApiError {
	err := r.DB.Refresh(&upCategory, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// Delete delete an existent category from DB.
func (r CategoryRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Category{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
