package repository

import (
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// BookRepository is responsible of getting information from DB.
type BookRepository struct {
	DB *gorm.DB
}

// Get returns all books from DB.
func (r BookRepository) Get() (books []model.Book, err error) {
	result := r.DB.Find(&books)
	if err = result.Error; err != nil {
		return nil, err
	}

	return
}

// Find return one book from DB by ID.
func (r BookRepository) Find(id int) (book model.Book, err error) {
	result := r.DB.Model(model.Book{}).Where("id = ?", id).First(&book)
	if err = result.Error; err != nil {
		return model.Book{}, err
	}

	return
}

// Create persist a book to the DB.
func (r BookRepository) Create(book model.Book) (uint, error) {
	result := r.DB.Create(&book)
	if err := result.Error; err != nil {
		return 0, err
	}

	return book.ID, nil
}

// Update update an existent book.
func (r BookRepository) Update(id int, upBook model.Book) (model.Book, error) {
	book, err := r.Find(id)
	if err != nil {
		return model.Book{}, err
	}

	result := r.DB.Model(&book).Updates(upBook)
	if err := result.Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// Delete delete an existent book from DB.
func (r BookRepository) Delete(id int) (err error) {
	book, err := r.Find(id)
	if err != nil {
		return
	}

	err = r.DB.Delete(&book).Error

	return
}
