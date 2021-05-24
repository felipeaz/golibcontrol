package repository

import (
	"fmt"
	"log"
	"net/http"

	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/book/model"
	categoryModel "github.com/FelipeAz/golibcontrol/internal/app/domain/category/model"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

type BookCategoryRepository struct {
	DB *gorm.DB
}

// GetCategoriesByIds returns categories by ids if exists on DB or an error
func (r BookCategoryRepository) GetCategoriesByIds(categoriesIds []uint) (categories []uint, apiError *errors.ApiError) {
	for _, categoryId := range categoriesIds {
		var category categoryModel.Category
		result := r.DB.First(&category, categoryId)
		if err := result.Error; err != nil {
			return nil, &errors.ApiError{
				Status: http.StatusBadRequest,
				Error:  fmt.Sprintf("category with ID: %v not found", categoryId),
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
		bookCategory := bookModel.BookCategory{
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
	result := r.DB.Delete(bookModel.BookCategory{}, "book_id = ?", bookId)
	if err := result.Error; err != nil {
		log.Println(err)
	}
}
