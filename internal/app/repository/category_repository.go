package repository

import (
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type CategoryRepository struct {
	DB *gorm.DB
}

// Get returns all categories.
func (r CategoryRepository) Get() (categories []model.Category, err error) {
	result := r.DB.Find(&categories)
	if err = result.Error; err != nil {
		return nil, err
	}

	return
}

// Find return one category from DB by ID.
func (r CategoryRepository) Find(id int) (category model.Category, err error) {
	result := r.DB.Model(model.Category{}).Where("id = ?", id).First(&category)
	if err = result.Error; err != nil {
		return model.Category{}, err
	}

	return
}

// Create persist a category to the DB.
func (r CategoryRepository) Create(category model.Category) (uint, error) {
	result := r.DB.Create(&category)
	if err := result.Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

// Update update an existent category.
func (r CategoryRepository) Update(id int, upCategory model.Category) (model.Category, error) {
	category, err := r.Find(id)
	if err != nil {
		return model.Category{}, err
	}

	result := r.DB.Model(&category).Updates(upCategory)
	if err := result.Error; err != nil {
		return model.Category{}, err
	}

	return category, nil
}

// Delete delete an existent category from DB.
func (r CategoryRepository) Delete(id int) (err error) {
	category, err := r.Find(id)
	if err != nil {
		return
	}

	err = r.DB.Delete(&category).Error
	return
}
