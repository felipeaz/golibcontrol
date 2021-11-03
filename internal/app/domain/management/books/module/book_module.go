package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// BookModule process the request recieved from handler.
type BookModule struct {
	Repository _interface.BookRepositoryInterface
	Log        logger.LogInterface
}

func NewBookModule(repo _interface.BookRepositoryInterface, log logger.LogInterface) BookModule {
	return BookModule{
		Repository: repo,
		Log:        log,
	}
}

// Get returns all books on DB.
func (m BookModule) Get() ([]model.Book, *errors.ApiError) {
	return m.Repository.Get()
}

// Find returns all books on DB.
func (m BookModule) Find(id string) (model.Book, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a book to the database.
func (m BookModule) Create(book model.Book) (uint, *errors.ApiError) {
	return m.Repository.Create(book)
}

// Update update an existent book.
func (m BookModule) Update(id string, upBook model.Book) *errors.ApiError {
	return m.Repository.Update(id, upBook)
}

// Delete delete an existent book.
func (m BookModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
