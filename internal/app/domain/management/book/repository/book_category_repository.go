package repository

import (
	"log"
	"net/http"
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
		_, err := r.DB.Fetch(&category, strconv.Itoa(int(categoryId)))
		if err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.UpdateFailMessage,
				Error:   err.Error(),
			}
		}
		categories = append(categories, category.ID)
	}
	return categories, nil
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
		err := r.DB.Persist(&bookCategory)
		if err != nil {
			log.Println(errors.FailedToCreateBookCategoryMessage, err.Error())
		}
	}
}

// DeleteCategories removes a Book categories from DB
func (r BookCategoryRepository) DeleteCategories(bookId uint) {
	err := r.DB.Remove(bookCategoryModel.BookCategory{}, strconv.Itoa(int(bookId)))
	if err != nil {
		log.Println(err.Error())
	}
}
