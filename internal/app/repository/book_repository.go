package repository

import (
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
)

// BookRepository is responsible of getting/saving information from DB.
type BookRepository struct {
	DB                     *gorm.DB
	BookCategoryRepository BookCategoryRepository
}

// Get returns all books from DB.
func (r BookRepository) Get() (books []model.Book, err error) {
	result := r.DB.Preload("BookCategory").Find(&books)
	if err = result.Error; err != nil {
		return nil, err
	}

	return
}

// Find return one book from DB by ID.
func (r BookRepository) Find(id int) (book model.Book, err error) {
	result := r.DB.Preload("BookCategory").Model(model.Book{}).Where("id = ?", id).First(&book)
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

	r.AfterCreate(book.ID, book.CategoriesId)
	return book.ID, nil
}

// Update update an existent book.
func (r BookRepository) Update(id int, upBook model.Book) (model.Book, error) {
	book, err := r.Find(id)
	if err != nil {
		return model.Book{}, err
	}

	r.BeforeUpdate(book.ID, upBook.CategoriesId)
	result := r.DB.Model(&book).Updates(upBook)
	if err = result.Error; err != nil {
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

	r.BeforeDelete(book.ID)
	err = r.DB.Delete(&book).Error
	return
}

// AfterCreate persists categories on BookCategory Table after the book persists.
func (r BookRepository) AfterCreate(bookId uint, categoriesId string) {
	r.BookCategoryRepository.CreateCategories(bookId, pkg.ExtractCategoryId(categoriesId))
}

// BeforeUpdate removes book category before updating data from DB.
func (r BookRepository) BeforeUpdate(bookId uint, categoriesId string) {
	r.BookCategoryRepository.DeleteCategories(bookId)
	r.BookCategoryRepository.CreateCategories(bookId, pkg.ExtractCategoryId(categoriesId))
}

// BeforeDelete removes book category before removing data from DB.
func (r BookRepository) BeforeDelete(bookId uint) {
	r.BookCategoryRepository.DeleteCategories(bookId)
}
