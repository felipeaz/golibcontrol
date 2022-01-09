package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

// CategoryModule process the request recieved from handler
type CategoryModule struct {
	Repository categories.Repository
	Log        logger.LogInterface
}

func NewCategoryModule(repo categories.Repository, log logger.LogInterface) CategoryModule {
	return CategoryModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all categories.
func (m CategoryModule) Get() ([]categories.Category, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one category by ID.
func (m CategoryModule) Find(id string) (categories.Category, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a category to the database.
func (m CategoryModule) Create(category categories.Category) (*categories.Category, *errors.ApiError) {
	return m.Repository.Create(category)
}

// Update update an existent category.
func (m CategoryModule) Update(id string, upCategory categories.Category) *errors.ApiError {
	return m.Repository.Update(id, upCategory)
}

// Delete delete an existent category.
func (m CategoryModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
