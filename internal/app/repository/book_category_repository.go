package repository

import (
	"log"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type BookCategoryRepository struct {
	DB *gorm.DB
}

// CreateCategories persists category on DB if exists.
func (r BookCategoryRepository) CreateCategories(bookId uint, categoriesIds []uint) {
	if len(categoriesIds) <= 0 {
		return
	}

	for _, categoryId := range categoriesIds {
		var category model.Category
		result := r.DB.First(&category, categoryId)
		if err := result.Error; err != nil {
			log.Println(err.Error())
			return
		}

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
