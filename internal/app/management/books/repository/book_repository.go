package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	registry "github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/app/filters"
	"github.com/FelipeAz/golibcontrol/internal/app/management/books/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
func (r BookRepository) Get() ([]books.Book, *errors.ApiError) {
	tx := r.DB.Preload("BookCategories", "BookCategories.Category", "Registry")
	result, err := r.DB.Find(tx, &[]books.Book{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceBookObj(result)
}

// Find return one book from DB by ID.
func (r BookRepository) Find(id string) (books.Book, *errors.ApiError) {
	tx := r.DB.Preload("BookCategories", "BookCategories.Category", "Registry")
	tx = r.DB.Where(tx, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &books.Book{})
	if err != nil {
		return books.Book{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToBookObj(result)
}

// GetByFilter return books from query string.
func (r BookRepository) GetByFilter(filter books.Filter) ([]books.Book, *errors.ApiError) {
	queryString := filters.BuildQueryFromFilter(filter)

	tx := r.DB.Preload("BookCategories", "BookCategories.Category", "Registry")
	if filter.Categories != "" {
		tx = r.DB.Join(tx, fmt.Sprintf("JOIN %s ON %s.book_id = %s.id",
			books.BookCategories{}.TableName(),
			books.BookCategories{}.TableName(),
			books.Book{}.TableName()))
	}
	if filter.RegistryNumber != "" {
		tx = r.DB.Join(tx, fmt.Sprintf("JOIN %s ON %s.book_id = %s.id",
			registry.Registry{}.TableName(),
			registry.Registry{}.TableName(),
			books.Book{}.TableName()))
	}
	tx = r.DB.Where(tx, queryString)
	tx = r.DB.Group(tx, fmt.Sprintf("%s.id", books.Book{}.TableName()))

	result, err := r.DB.Find(tx, &[]books.Book{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceBookObj(result)
}

// Create persist a book to the DB.
func (r BookRepository) Create(book books.Book) (*books.Book, *errors.ApiError) {
	err := r.DB.Persist(&book)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &book, nil
}

// Update update an existent book.
func (r BookRepository) Update(id string, upBook books.Book) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upBook)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	err = r.DB.Set(tx, &upBook, "available", upBook.Available)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &books.Book{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
