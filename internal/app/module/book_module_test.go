package module

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/mock"
)

func TestGetBook(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Get()
	book := books[0]

	// Validation
	assert.Nil(t, apiError)
	assert.Equal(t, 25, int(book.ID))
	assert.Equal(t, "Mocked Book", book.Title)
	assert.Equal(t, "Mocked Author", book.Author)
	assert.Equal(t, "123", book.RegisterNumber)
	assert.Equal(t, true, book.Available)
}

func TestGetBookError(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestError = true
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Get()

	// Validation
	assert.NotNil(t, apiError)
	assert.Nil(t, books)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked test get error", apiError.Error)
}

func TestFindBook(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	book, apiError := m.Find("25")

	// Validation
	assert.Nil(t, apiError)
	assert.NotEqual(t, model.Book{}, book)
	assert.Equal(t, 25, int(book.ID))
	assert.Equal(t, "Mocked Book", book.Title)
	assert.Equal(t, "Mocked Book", book.Title)
	assert.Equal(t, "Mocked Author", book.Author)
	assert.Equal(t, "123", book.RegisterNumber)
	assert.Equal(t, true, book.Available)
}

func TestFindBookError(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestError = true
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Find("25")

	// Validation

	assert.NotNil(t, apiError)
	assert.Equal(t, model.Book{}, books)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked test find error", apiError.Error)
}

func TestFindBookNotFoundError(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestNotFoundError = true
	m := BookModule{BookRepository: bookRepositoryMock}

	// Exec
	books, apiError := m.Find("25")

	// Validation

	assert.NotNil(t, apiError)
	assert.Equal(t, model.Book{}, books)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked test find not found error", apiError.Error)
}

func TestCreateBook(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	m := BookModule{BookRepository: bookRepositoryMock}
	book := model.Book{
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		Available:      true,
	}

	// Exec
	bookId, apiError := m.Create(book)

	// Validation
	assert.Nil(t, apiError)
	assert.NotNil(t, bookId)
	assert.Equal(t, 25, int(bookId))
}

func TestCreateBookWithCategoryError(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestCategoryNotFoundError = true
	m := BookModule{BookRepository: bookRepositoryMock}
	book := model.Book{
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		CategoriesId:   "1,2,5",
		Available:      true,
	}

	// Exec
	bookId, apiError := m.Create(book)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(bookId))
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "category with ID: 5 not found", apiError.Error)
}

func TestCreateBookWithError(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestError = true
	m := BookModule{BookRepository: bookRepositoryMock}
	book := model.Book{
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		CategoriesId:   "1,2,5",
		Available:      true,
	}

	// Exec
	bookId, apiError := m.Create(book)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(bookId))
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestBookUpdate(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	m := BookModule{BookRepository: bookRepositoryMock}
	id := "25"
	upBook := model.Book{
		RegisterNumber: "123",
		Title:          "Mocked Book Updated",
		Author:         "Mocked Author Updated",
		Available:      false,
	}

	// Exec
	book, apiError := m.Update(id, upBook)

	// Validation
	assert.Nil(t, apiError)
	assert.NotEqual(t, model.Book{}, book)
	assert.Equal(t, 25, int(book.ID))
	assert.Equal(t, upBook.Title, book.Title)
	assert.Equal(t, upBook.Author, book.Author)
	assert.Equal(t, upBook.Available, book.Available)
}

func TestUpdateBookNotFound(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestNotFoundError = true
	m := BookModule{BookRepository: bookRepositoryMock}
	id := "25"
	book := model.Book{
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		Available:      true,
	}

	// Exec
	book, apiError := m.Update(id, book)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "book not found", apiError.Error)
}

func TestUpdateBookCategoryNotFound(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestCategoryNotFoundError = true
	m := BookModule{BookRepository: bookRepositoryMock}
	id := "25"
	book := model.Book{
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		CategoriesId:   "1,2,5",
		Available:      true,
	}

	// Exec
	book, apiError := m.Update(id, book)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "category with ID: 5 not found", apiError.Error)
}

func TestUpdateBookWithError(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestError = true
	m := BookModule{BookRepository: bookRepositoryMock}
	id := "25"
	book := model.Book{
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		CategoriesId:   "1,2,5",
		Available:      true,
	}

	// Exec
	book, apiError := m.Update(id, book)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestDeleteBook(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	m := BookModule{BookRepository: bookRepositoryMock}
	id := "25"

	// Exec
	apiError := m.Delete(id)

	// Validation
	assert.Nil(t, apiError)
}

func TestDeleteBookNotFound(t *testing.T) {
	// Init
	var bookRepositoryMock = mock.BookRepositoryMock{}
	bookRepositoryMock.TestNotFoundError = true
	m := BookModule{BookRepository: bookRepositoryMock}
	id := "25"

	// Exec
	apiError := m.Delete(id)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "book not found", apiError.Error)
}
