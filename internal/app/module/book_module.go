package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/interfaces/repository"
)

// BookModule process the request recieved from handler.
type BookModule struct {
	BookRepository repository.BookRepositoryInterface
}

// Get returns all books on DB.
func (m BookModule) Get() (books []model.Book, apiError *errors.ApiError) {
	books, apiError = m.BookRepository.Get()
	return
}

// Find returns all books on DB.
func (m BookModule) Find(id string) (book model.Book, apiError *errors.ApiError) {
	book, apiError = m.BookRepository.Find(id)
	return
}

// Create persist a book to the database.
func (m BookModule) Create(book model.Book) (id uint, apiError *errors.ApiError) {
	id, apiError = m.BookRepository.Create(book)
	return
}

// Update update an existent book.
func (m BookModule) Update(id string, upBook model.Book) (book model.Book, apiError *errors.ApiError) {
	book, apiError = m.BookRepository.Update(id, upBook)
	return
}

// Delete delete an existent book.
func (m BookModule) Delete(id string) (apiError *errors.ApiError) {
	apiError = m.BookRepository.Delete(id)
	return
}
