package module

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type BookRepositoryMock struct {
	TestError         bool
	TestNotFoundError bool
}

func (r BookRepositoryMock) Get() (books []model.Book, apiError *errors.ApiError) {
	if r.TestError {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked test get error",
		}
	}

	books = []model.Book{
		model.Book{
			RegisterNumber: "123",
			Title:          "Mocked Book",
			Author:         "Mocked Author",
			Available:      true,
		},
	}

	return
}

func (r BookRepositoryMock) Find(id string) (book model.Book, apiError *errors.ApiError) {
	if r.TestError {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked test find error",
		}
	} else if r.TestNotFoundError {

		return model.Book{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "mocked test find not found error",
		}
	}

	book = model.Book{
		ID:             uint(25),
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		Available:      false,
	}

	return
}

func (r BookRepositoryMock) Create(book model.Book) (uint, *errors.ApiError) {
	return 0, nil
}

func (r BookRepositoryMock) Update(id string, upBook model.Book) (model.Book, *errors.ApiError) {
	return model.Book{}, nil
}

func (r BookRepositoryMock) Delete(id string) (apiError *errors.ApiError) {
	return nil
}

func (r BookRepositoryMock) BeforeCreate(categoriesId string) ([]uint, *errors.ApiError) {
	return []uint{}, nil
}

func (r BookRepositoryMock) AfterCreate(bookId uint, categoriesId []uint) {}

func (r BookRepositoryMock) BeforeUpdate(bookId uint, categoriesId string) *errors.ApiError {
	return nil
}

func (r BookRepositoryMock) BeforeDelete(bookId uint) {}

func TestGet(t *testing.T) {
	// Init
	var bookRepositoryMock = BookRepositoryMock{}
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Get()
	book := books[0]

	// Validation
	assert.Nil(t, apiError)
	assert.Equal(t, "Mocked Book", book.Title)
	assert.Equal(t, "Mocked Author", book.Author)
	assert.Equal(t, "123", book.RegisterNumber)
	assert.Equal(t, true, book.Available)
}

func TestGetError(t *testing.T) {
	// Init
	var bookRepositoryMock = BookRepositoryMock{}
	bookRepositoryMock.TestError = true
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Get()

	// Validation
	assert.Nil(t, books)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked test get error", apiError.Error)
}

func TestFind(t *testing.T) {
	// Init
	var bookRepositoryMock = BookRepositoryMock{}
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	book, apiError := m.Find("25")

	// Validation
	assert.Nil(t, apiError)
	assert.Equal(t, 25, int(book.ID))
	assert.Equal(t, "Mocked Book", book.Title)
	assert.Equal(t, "Mocked Book", book.Title)
	assert.Equal(t, "Mocked Author", book.Author)
	assert.Equal(t, "123", book.RegisterNumber)
	assert.Equal(t, false, book.Available)
}

func TestFindError(t *testing.T) {
	// Init
	var bookRepositoryMock = BookRepositoryMock{}
	bookRepositoryMock.TestError = true
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Find("25")

	// Validation

	assert.Equal(t, model.Book{}, books)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked test find error", apiError.Error)
}

func TestFindNotFoundError(t *testing.T) {
	// Init
	var bookRepositoryMock = BookRepositoryMock{}
	bookRepositoryMock.TestNotFoundError = true
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Find("25")

	// Validation

	assert.Equal(t, model.Book{}, books)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked test find not found error", apiError.Error)
}

func TestCreate(t *testing.T) {
	// Init

}
