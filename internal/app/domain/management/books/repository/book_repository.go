package repository

import (
	"fmt"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model/converter"
)

// BookRepository is responsible for getting/saving information from DB.
type BookRepository struct {
	DB database.GORMServiceInterface
}

func NewBookRepository(db database.GORMServiceInterface) BookRepository {
	return BookRepository{
		DB: db,
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
	return converter.ConvertToSliceBookObj(result)
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
	return converter.ConvertToBookObj(result)
}

// GetWhere return books from query string.
func (r BookRepository) GetWhere(queryBook model.QueryBook) ([]model.Book, *errors.ApiError) {
	join := fmt.Sprintf("JOIN book_categories ON book_categories.book_id = books.id")
	query := fmt.Sprintf("book_categories.category_id IN (%s)", *queryBook.Categories)
	result, err := r.DB.FetchAllWithQueryAndPreload(&[]model.Book{}, query, "BookCategory", join)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return converter.ConvertToSliceBookObj(result)
}

// Create persist a book to the DB.
func (r BookRepository) Create(book model.Book) (uint, *errors.ApiError) {
	err := r.DB.Persist(&book)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return book.ID, nil
}

// Update update an existent book.
func (r BookRepository) Update(id string, upBook model.Book) *errors.ApiError {
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
