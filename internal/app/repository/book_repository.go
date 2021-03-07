package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"gorm.io/gorm"
)

// BookRepository is responsible of getting information from DB.
type BookRepository struct {
	DB *gorm.DB
}

// GetBooks returns all books on DB.
func (r BookRepository) GetBooks() (books []model.Book, err error){
	result := r.DB.Find(&books)
	if err = result.Error; err != nil {
		return nil, err
	}

	return
}

// GetBook return one book from DB by ID.
func (r BookRepository) GetBook(id int) (book model.Book, err error) {
	result := r.DB.Model(&model.Book{}).Where("id = ?", id).First(&book)
	if err = result.Error; err != nil {
		return model.Book{}, err
	}

	return
}

// CreateBook persist a book to the database.
func (r BookRepository) CreateBook(book model.Book) (uint, error) {
	result := r.DB.Create(&book)
	if err := result.Error; err != nil {
		return 0, err
	}

	return book.ID, nil
}

// UpdateBook update an existent book.
func (r BookRepository) UpdateBook(id int, upBook model.Book) (book model.Book, err error) {
	result := r.DB.Model(model.Book{}).Updates(upBook)
	if err = result.Error; err != nil {
		return model.Book{}, err
	}

	return
}

func (r BookRepository) DeleteBook(id int) error {
	result := r.DB.Delete(&model.Book{}, id)

	return result.Error
}
