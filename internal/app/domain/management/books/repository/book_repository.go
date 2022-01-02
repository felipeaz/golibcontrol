package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/pkg"
	"reflect"
	"strings"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
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
	result, err := r.DB.FetchAllWithPreload(&[]books.Book{}, "Category")
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
	result, err := r.DB.FetchWithPreload(&books.Book{}, id, "Category")
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
	queryString := r.buildQueryFromFilter(filter)
	join := fmt.Sprintf("JOIN book_categories ON book_categories.book_id = books.id")
	result, err := r.DB.FetchAllWithQueryAndPreload(
		&[]books.Book{},
		queryString,
		"Category",
		join,
		"books.id",
	)
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
	err := r.DB.Refresh(&upBook, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	err = r.DB.Set(&upBook, id, "available", upBook.Available)
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
	err := r.DB.Remove(&books.Book{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r BookRepository) buildQueryFromFilter(filter books.Filter) string {
	var query []string

	// reflect allows accessing type metadata (ex: struct tags)
	fields := reflect.TypeOf(filter)
	for _, name := range filter.GetFieldNames() {
		field, ok := fields.FieldByName(name)
		if !ok {
			continue
		}

		fieldValue := reflect.ValueOf(filter).FieldByName(name)
		if !fieldValue.IsZero() {
			var qs string

			switch field.Tag.Get("array") {
			case "false":
				qs = fmt.Sprintf("%s = %v", field.Tag.Get("column"), fieldValue.Interface())
			default:
				qs = fmt.Sprintf("%s IN (%v)", field.Tag.Get("column"), fieldValue.Interface())
			}

			query = append(query, qs)
		}
	}

	return strings.Join(query, " AND ")
}
