package repository

import (
	"fmt"
	"reflect"
	"strings"

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

// GetByFilter return books from query string.
func (r BookRepository) GetByFilter(filter model.Filter) ([]model.Book, *errors.ApiError) {
	queryString := r.buildQueryFromFilter(filter)
	join := fmt.Sprintf("JOIN book_categories ON book_categories.book_id = books.id")
	result, err := r.DB.FetchAllWithQueryAndPreload(
		&[]model.Book{},
		queryString,
		"BookCategory",
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

func (r BookRepository) buildQueryFromFilter(filter model.Filter) string {
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
