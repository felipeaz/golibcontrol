package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// CategoryModule process the request recieved from handler
type CategoryModule struct {
	Repository _interface.CategoryRepositoryInterface
	Log        logger.LogInterface
}

func NewCategoryModule(repo _interface.CategoryRepositoryInterface, log logger.LogInterface) CategoryModule {
	return CategoryModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all categories.
func (m CategoryModule) Get() ([]model.Category, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one category by ID.
func (m CategoryModule) Find(id string) (model.Category, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a category to the database.
func (m CategoryModule) Create(category model.Category) (*model.Category, *errors.ApiError) {
	return m.Repository.Create(category)
}

// Update update an existent category.
func (m CategoryModule) Update(id string, upCategory model.Category) *errors.ApiError {
	return m.Repository.Update(id, upCategory)
}

// Delete delete an existent category.
func (m CategoryModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
