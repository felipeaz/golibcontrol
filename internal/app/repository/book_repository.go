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
func (r BookRepository) Get() (books []model.Book, err error){
	result := r.DB.Find(&books)
	if err = result.Error; err != nil {
		return nil, err
	}

	return
}

// GetBook return one book from DB by ID.
func (r BookRepository) Find(id int) (book model.Book, err error) {
	result := r.DB.Model(&model.Book{}).Where("id = ?", id).First(&book)
	if err = result.Error; err != nil {
		return model.Book{}, err
	}

	return
}

// CreateBook persist a book to the database.
func (r BookRepository) Create(book model.Book) (uint, error) {
	result := r.DB.Create(&book)
	if err := result.Error; err != nil {
		return 0, err
	}

	return book.ID, nil
}

// UpdateBook update an existent book.
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

func (r BookRepository) Delete(id int) error {
	book, err := r.Find(id)
	if err != nil {
		return err
	}

	result := r.DB.Delete(&book)

	return result.Error
}
