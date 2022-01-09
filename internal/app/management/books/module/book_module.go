package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	_pkg "github.com/FelipeAz/golibcontrol/internal/app/domain/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/management/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

// BookModule process the request received from handler.
type BookModule struct {
	Repository             books.Repository
	BookCategoryRepository books.CategoryRepository
	Log                    logger.LogInterface
}

func NewBookModule(
	repo books.Repository,
	cRepo books.CategoryRepository,
	log logger.LogInterface) BookModule {
	return BookModule{
		Repository:             repo,
		BookCategoryRepository: cRepo,
		Log:                    log,
	}
}

// Get returns all books on DB.
func (m BookModule) Get() ([]books.Book, *errors.ApiError) {
	return m.Repository.Get()
}

// GetByFilter return all books from Query
func (m BookModule) GetByFilter(filter books.Filter) ([]books.Book, *errors.ApiError) {
	return m.Repository.GetByFilter(filter)
}

// Find returns all books on DB.
func (m BookModule) Find(id string) (books.Book, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a book to the database.
func (m BookModule) Create(book books.Book) (*books.Book, *errors.ApiError) {
	categoriesIds := pkg.ExtractCategoryId(book.CategoriesId)
	resp, apiError := m.Repository.Create(book)
	if apiError != nil {
		return nil, apiError
	}
	m.setBookCategories(resp.ID, categoriesIds)
	return resp, nil
}

// Update update an existent book.
func (m BookModule) Update(id string, upBook books.Book) *errors.ApiError {
	categoriesIds := pkg.ExtractCategoryId(upBook.CategoriesId)
	apiError := m.Repository.Update(id, upBook)
	if apiError != nil {
		return apiError
	}
	bookId, apiError := _pkg.ParseStringToId(id)
	if apiError != nil {
		return apiError
	}
	m.unsetBookCategories(bookId)
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
	m.BookCategoryRepository.CreateCategories(bookId, categoriesId)
}

func (m BookModule) unsetBookCategories(bookId uint) {
	m.BookCategoryRepository.DeleteCategories(bookId)
}
