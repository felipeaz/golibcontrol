package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/category/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/category/repository/interface"
)

// CategoryModule process the request recieved from handler
type CategoryModule struct {
	Repository _interface.CategoryRepositoryInterface
}

// Get returns all categories.
func (m CategoryModule) Get() (categories []model.Category, apiError *errors.ApiError) {
	categories, apiError = m.Repository.Get()
	return
}

// Find return one category by ID.
func (m CategoryModule) Find(id string) (category model.Category, apiError *errors.ApiError) {
	category, apiError = m.Repository.Find(id)
	return
}

// Create persist a category to the database.
func (m CategoryModule) Create(category model.Category) (id uint, apiError *errors.ApiError) {
	id, apiError = m.Repository.Create(category)
	return
}

// Update update an existent category.
func (m CategoryModule) Update(id string, upCategory model.Category) (category model.Category, apiError *errors.ApiError) {
	category, apiError = m.Repository.Update(id, upCategory)
	return
}

// Delete delete an existent category.
func (m CategoryModule) Delete(id string) (apiError *errors.ApiError) {
	apiError = m.Repository.Delete(id)
	return
}
