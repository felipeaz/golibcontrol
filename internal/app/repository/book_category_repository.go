package repository

import (
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type BookCategoryRepository struct {
	DB *gorm.DB
}

// ValidateCategories validate if categories exists on DB
func (r BookCategoryRepository) ValidateCategories(categoriesIds []uint) (categories []uint, apiError *errors.ApiError) {
	for _, categoryId := range categoriesIds {
		var category model.Category
		result := r.DB.First(&category, categoryId)
		if err := result.Error; err != nil {
			return nil, &errors.ApiError{
				Status: http.StatusBadRequest,
				Error:  "category not found",
			}
		}

		categories = append(categories, category.ID)
	}

	return
}

// CreateCategories persists category on DB if exists.
func (r BookCategoryRepository) CreateCategories(bookId uint, categoriesIds []uint) {
	if len(categoriesIds) <= 0 {
		return
	}

	for _, categoryId := range categoriesIds {
		bookCategory := model.BookCategory{
			BookID:     bookId,
			CategoryID: categoryId,
		}

		err := r.DB.Create(&bookCategory).Error
		if err != nil {
			log.Println(err)
		}
	}
}

// DeleteCategories removes a Book categories from DB
func (r BookCategoryRepository) DeleteCategories(bookId uint) {
	result := r.DB.Delete(model.BookCategory{}, "book_id = ?", bookId)
	if err := result.Error; err != nil {
		log.Println(err)
	}
}
