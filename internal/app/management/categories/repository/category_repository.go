package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/app/management/categories/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
	result, err := r.DB.Find(nil, &[]categories.Category{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceCategoryObj(result)
}

// Find return one category from DB by ID.
func (r CategoryRepository) Find(id string) (categories.Category, *errors.ApiError) {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &categories.Category{})
	if err != nil {
		return categories.Category{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToCategoryObj(result)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upCategory)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// Delete an existent category from DB.
func (r CategoryRepository) Delete(id string) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &categories.Category{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
