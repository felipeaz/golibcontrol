package repository

import (
	"log"
	"strconv"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	bookCategoryModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/model"
	categoryModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/model"
)

type BookCategoryRepository struct {
	DB database.GORMServiceInterface
}

func NewBookCategoryRepository(db database.GORMServiceInterface) BookCategoryRepository {
	return BookCategoryRepository{
		DB: db,
	}
}

// GetCategoriesByIds returns categories by ids if exists on DB or an error
func (r BookCategoryRepository) GetCategoriesByIds(categoriesIds []uint) (categories []uint, apiError *errors.ApiError) {
	for _, categoryId := range categoriesIds {
		var category categoryModel.Category
		_, err := r.DB.Find(&category, strconv.Itoa(int(categoryId)))
		if err != nil {
			return nil, err
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
		bookCategory := bookCategoryModel.BookCategory{
			BookID:     bookId,
			CategoryID: categoryId,
		}
		apiError := r.DB.Create(&bookCategory)
		if apiError != nil {
			log.Println(errors.FailedToCreateBookCategoryMessage, apiError.Error)
		}
	}
}

// DeleteCategories removes a Book categories from DB
func (r BookCategoryRepository) DeleteCategories(bookId uint) {
	apiErr := r.DB.Delete(bookCategoryModel.BookCategory{}, strconv.Itoa(int(bookId)))
	if apiErr != nil {
		log.Println(apiErr)
	}
}
