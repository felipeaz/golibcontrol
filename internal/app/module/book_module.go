package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// BookModule process the request recieved from handler.
type BookModule struct {
	Repository repository.BookRepository
}

// Get returns all books on DB.
func (m BookModule) Get() (books []model.Book, err error) {
	books, err = m.Repository.Get()
	return
}

// Find returns all books on DB.
func (m BookModule) Find(id int) (book model.Book, err error) {
	book, err = m.Repository.Find(id)
	return
}

// Create persist a book to the database.
func (m BookModule) Create(book model.Book) (id uint, err error) {
	id, err = m.Repository.Create(book)
	return
}

// Update update an existent book.
func (m BookModule) Update(id int, upBook model.Book) (book model.Book, err error) {
	book, err = m.Repository.Update(id, upBook)
	return
}

// Delete delete an existent book.
func (m BookModule) Delete(id int) (err error) {
	err = m.Repository.Delete(id)
	return
}
