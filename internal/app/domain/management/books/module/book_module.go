package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/pkg"
	_pkg "github.com/FelipeAz/golibcontrol/internal/app/domain/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// BookModule process the request received from handler.
type BookModule struct {
	Repository             _interface.BookRepositoryInterface
	BookCategoryRepository _interface.BookCategoryRepositoryInterface
	Log                    logger.LogInterface
}

func NewBookModule(
	repo _interface.BookRepositoryInterface,
	cRepo _interface.BookCategoryRepositoryInterface,
	log logger.LogInterface) BookModule {
	return BookModule{
		Repository:             repo,
		BookCategoryRepository: cRepo,
		Log:                    log,
	}
}

// Get returns all books on DB.
func (m BookModule) Get() ([]model.Book, *errors.ApiError) {
	return m.Repository.Get()
}

// GetWhere return all books from Query
func (m BookModule) GetWhere(queryBook model.QueryBook) ([]model.Book, *errors.ApiError) {
	return m.Repository.GetWhere(queryBook)
}

// Find returns all books on DB.
func (m BookModule) Find(id string) (model.Book, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a book to the database.
func (m BookModule) Create(book model.Book) (uint, *errors.ApiError) {
	categoriesIds := pkg.ExtractCategoryId(book.CategoriesId)
	bookId, apiError := m.Repository.Create(book)
	if apiError != nil {
		return 0, apiError
	}
	m.setBookCategories(bookId, categoriesIds)
	return bookId, nil
}

// Update update an existent book.
func (m BookModule) Update(id string, upBook model.Book) *errors.ApiError {
	categoriesIds := pkg.ExtractCategoryId(upBook.CategoriesId)
	apiError := m.Repository.Update(id, upBook)
	if apiError != nil {
		return apiError
	}
	bookId, apiError := _pkg.ParseStringToId(id)
	if apiError != nil {
		return apiError
	}
	m.setBookCategories(bookId, categoriesIds)
	return nil
}

// Delete delete an existent book.
func (m BookModule) Delete(id string) *errors.ApiError {
	parsedId, apiError := _pkg.ParseStringToId(id)
	if apiError != nil {
		return apiError
	}
	m.unsetBookCategories(parsedId)
	return m.Repository.Delete(id)
}

func (m BookModule) setBookCategories(bookId uint, categoriesId []uint) {
	m.unsetBookCategories(bookId)
	m.BookCategoryRepository.CreateCategories(bookId, categoriesId)
}

func (m BookModule) unsetBookCategories(bookId uint) {
	m.BookCategoryRepository.DeleteCategories(bookId)
}
