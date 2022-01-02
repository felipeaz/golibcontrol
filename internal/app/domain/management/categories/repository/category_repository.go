package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories/pkg"
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
func (r CategoryRepository) Get() ([]categories.Category, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]categories.Category{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	categories, apiError := pkg.ParseToSliceCategoryObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return categories, nil
}

// Find return one category from DB by ID.
func (r CategoryRepository) Find(id string) (categories.Category, *errors.ApiError) {
	result, err := r.DB.Fetch(&categories.Category{}, id)
	if err != nil {
		return categories.Category{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	category, apiError := pkg.ParseToCategoryObj(result)
	if apiError != nil {
		return categories.Category{}, apiError
	}

	return category, nil
}

// Create persist a category to the DB.
func (r CategoryRepository) Create(category categories.Category) (*categories.Category, *errors.ApiError) {
	err := r.DB.Persist(&category)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &category, nil
}

// Update update an existent category.
func (r CategoryRepository) Update(id string, upCategory categories.Category) *errors.ApiError {
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
	err := r.DB.Remove(&categories.Category{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
