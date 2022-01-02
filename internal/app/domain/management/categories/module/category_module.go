package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	domain "github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// CategoryModule process the request recieved from handler
type CategoryModule struct {
	Repository domain.Repository
	Log        logger.LogInterface
}

func NewCategoryModule(repo domain.Repository, log logger.LogInterface) CategoryModule {
	return CategoryModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all categories.
func (m CategoryModule) Get() ([]domain.Category, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one category by ID.
func (m CategoryModule) Find(id string) (domain.Category, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a category to the database.
func (m CategoryModule) Create(category domain.Category) (*domain.Category, *errors.ApiError) {
	return m.Repository.Create(category)
}

// Update update an existent category.
func (m CategoryModule) Update(id string, upCategory domain.Category) *errors.ApiError {
	return m.Repository.Update(id, upCategory)
}

// Delete delete an existent category.
func (m CategoryModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
