package repository

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
)

// BookRepository is responsible of getting/saving information from DB.
type BookRepository struct {
	DB                     *gorm.DB
	BookCategoryRepository BookCategoryRepository
}

// Get returns all books from DB.
func (r BookRepository) Get() (books []model.Book, apiError *errors.ApiError) {
	result := r.DB.Preload("BookCategory").Find(&books)
	if err := result.Error; err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.NotFoundMessage,
			Error:   err.Error(),
		}
	}

	return
}

// Find return one book from DB by ID.
func (r BookRepository) Find(id string) (book model.Book, apiError *errors.ApiError) {
	result := r.DB.Preload("BookCategory").Model(model.Book{}).Where("id = ?", id).First(&book)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return model.Book{}, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.NotFoundMessage,
				Error:   err.Error(),
			}
		}

		return model.Book{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.NotFoundMessage,
			Error:   "book not found",
		}
	}

	return
}

// Create persist a book to the DB.
func (r BookRepository) Create(book model.Book) (uint, *errors.ApiError) {
	result := r.DB.Create(&book)
	if err := result.Error; err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailedMessage,
			Error:   err.Error(),
		}
	}

	r.AfterCreate(book.ID, book.CategoriesId)
	return book.ID, nil
}

// Update update an existent book.
func (r BookRepository) Update(id string, upBook model.Book) (model.Book, *errors.ApiError) {
	book, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailedMessage
		return model.Book{}, apiError
	}

	r.BeforeUpdate(book.ID, upBook.CategoriesId)
	result := r.DB.Model(&book).Updates(upBook)
	if err := result.Error; err != nil {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailedMessage,
			Error:   err.Error(),
		}
	}

	return book, nil
}

// Delete delete an existent book from DB.
func (r BookRepository) Delete(id string) (apiError *errors.ApiError) {
	book, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailedMessage
		return
	}

	r.BeforeDelete(book.ID)
	err := r.DB.Delete(&book).Error
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailedMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AfterCreate persists categories on BookCategory Table after the book persists.
func (r BookRepository) AfterCreate(bookId uint, categoriesId string) {
	r.BookCategoryRepository.CreateCategories(bookId, pkg.ExtractCategoryId(categoriesId))
}

// BeforeUpdate removes book category before updating data from DB.
func (r BookRepository) BeforeUpdate(bookId uint, categoriesId string) {
	r.BookCategoryRepository.DeleteCategories(bookId)
	r.BookCategoryRepository.CreateCategories(bookId, pkg.ExtractCategoryId(categoriesId))
}

// BeforeDelete removes book category before removing data from DB.
func (r BookRepository) BeforeDelete(bookId uint) {
	r.BookCategoryRepository.DeleteCategories(bookId)
}
