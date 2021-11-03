package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model/converter"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/pkg"
	_pkg "github.com/FelipeAz/golibcontrol/internal/app/domain/pkg"
)

// BookRepository is responsible for getting/saving information from DB.
type BookRepository struct {
	DB                     database.GORMServiceInterface
	BookCategoryRepository _interface.BookCategoryRepositoryInterface
}

func NewBookRepository(db database.GORMServiceInterface, repo _interface.BookCategoryRepositoryInterface) BookRepository {
	return BookRepository{
		DB:                     db,
		BookCategoryRepository: repo,
	}
}

// Get returns all books from DB.
func (r BookRepository) Get() ([]model.Book, *errors.ApiError) {
	result, err := r.DB.FetchAllWithPreload(&[]model.Book{}, "BookCategory")
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	books, apiError := converter.ConvertToSliceBookObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return books, nil
}

// Find return one book from DB by ID.
func (r BookRepository) Find(id string) (model.Book, *errors.ApiError) {
	result, err := r.DB.FetchWithPreload(&model.Book{}, id, "BookCategory")
	if err != nil {
		return model.Book{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	book, apiError := converter.ConvertToBookObj(result)
	if apiError != nil {
		return model.Book{}, apiError
	}
	return book, nil
}

// Create persist a book to the DB.
func (r BookRepository) Create(book model.Book) (uint, *errors.ApiError) {
	categoriesIds, apiError := r.BeforeCreate(book.CategoriesId)
	if apiError != nil {
		apiError.Message = errors.CreateFailMessage
		if apiError.Error == errors.ItemNotFoundError {
			apiError.Error = errors.CategoryNotFoundError
		}
		return 0, apiError
	}

	err := r.DB.Persist(&book)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	r.AfterCreate(book.ID, categoriesIds)
	return book.ID, nil
}

// Update update an existent book.
func (r BookRepository) Update(id string, upBook model.Book) *errors.ApiError {
	parsedId, apiError := _pkg.ParseStringToId(id)
	if apiError != nil {
		return apiError
	}

	apiError = r.BeforeUpdate(parsedId, upBook.CategoriesId)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return apiError
	}
	err := r.DB.Refresh(&upBook, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// Delete delete an existent book from DB.
func (r BookRepository) Delete(id string) *errors.ApiError {
	parsedId, apiError := _pkg.ParseStringToId(id)
	if apiError != nil {
		return apiError
	}
	r.BeforeDelete(parsedId)
	err := r.DB.Remove(&model.Book{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// BeforeCreate validate if the request categories exists
func (r BookRepository) BeforeCreate(categoriesId string) ([]uint, *errors.ApiError) {
	return r.BookCategoryRepository.GetCategoriesByIds(pkg.ExtractCategoryId(categoriesId))
}

// AfterCreate persists categories on BookCategory Table after the book persists.
func (r BookRepository) AfterCreate(bookId uint, categoriesId []uint) {
	r.BookCategoryRepository.CreateCategories(bookId, categoriesId)
}

// BeforeUpdate removes book category before updating data from DB.
func (r BookRepository) BeforeUpdate(bookId uint, categoriesId string) *errors.ApiError {
	categoriesIdSlice, apiError := r.BeforeCreate(categoriesId)
	if apiError != nil {
		return apiError
	}
	r.BookCategoryRepository.DeleteCategories(bookId)
	r.BookCategoryRepository.CreateCategories(bookId, categoriesIdSlice)
	return nil
}

// BeforeDelete removes book category before removing data from DB.
func (r BookRepository) BeforeDelete(bookId uint) {
	r.BookCategoryRepository.DeleteCategories(bookId)
}
