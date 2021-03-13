package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// CategoryModule process the request recieved from handler
type CategoryModule struct {
	Repository repository.CategoryRepository
}

// Get returns all categories.
func (m CategoryModule) Get() (categories []model.Category, err error) {
	categories, err = m.Repository.Get()
	return
}

// Find return one category by ID.
func (m CategoryModule) Find(id int) (category model.Category, err error) {
	category, err = m.Repository.Find(id)
	return
}

// Create persist a category to the database.
func (m CategoryModule) Create(category model.Category) (id uint, err error) {
	id, err = m.Repository.Create(category)
	return
}

// Update update an existent category.
func (m CategoryModule) Update(id int, upCategory model.Category) (category model.Category, err error) {
	category, err = m.Repository.Update(id, upCategory)
	return
}

// Delete delete an existent category.
func (m CategoryModule) Delete(id int) (err error) {
	err = m.Repository.Delete(id)
	return
}
