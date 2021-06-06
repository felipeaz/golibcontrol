package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/book/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/book/model/converter"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/book/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
)

// BookRepository is responsible of getting/saving information from DB.
type BookRepository struct {
	DB                     database.GORMServiceInterface
	BookCategoryRepository _interface.BookCategoryRepositoryInterface
}

// Get returns all books from DB.
func (r BookRepository) Get() ([]model.Book, *errors.ApiError) {
	result, apiError := r.DB.GetWithPreload(&[]model.Book{}, "BookCategory")
	if apiError != nil {
		return nil, apiError
	}
	books, apiError := converter.ConvertToSliceBookObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return books, nil
}

// Find return one book from DB by ID.
func (r BookRepository) Find(id string) (model.Book, *errors.ApiError) {
	result, apiError := r.DB.FindWithPreload(&model.Book{}, id, "BookCategory")
	if apiError != nil {
		return model.Book{}, apiError
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

	apiError = r.DB.Create(&book)
	if apiError != nil {
		return 0, apiError
	}

	r.AfterCreate(book.ID, categoriesIds)
	return book.ID, nil
}

// Update update an existent book.
func (r BookRepository) Update(id string, upBook model.Book) *errors.ApiError {
	parsedId, apiError := pkg.ParseStringToId(id)
	if apiError != nil {
		return apiError
	}

	apiError = r.BeforeUpdate(parsedId, upBook.CategoriesId)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return apiError
	}
	return r.DB.Update(&upBook, id)
}

// Delete delete an existent book from DB.
func (r BookRepository) Delete(id string) *errors.ApiError {
	parsedId, err := pkg.ParseStringToId(id)
	if err != nil {
		return err
	}
	r.BeforeDelete(parsedId)
	return r.DB.Delete(&model.Book{}, id)
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
