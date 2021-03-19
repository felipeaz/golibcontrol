package mock

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type BookRepositoryMock struct {
	TestError                 bool
	TestCategoryNotFoundError bool
	TestNotFoundError         bool
}

func (r BookRepositoryMock) Get() (books []model.Book, apiError *errors.ApiError) {
	if r.TestError {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked test get error",
		}
	}

	books = []model.Book{
		model.Book{
			RegisterNumber: "123",
			Title:          "Mocked Book",
			Author:         "Mocked Author",
			Available:      true,
		},
	}

	return
}

func (r BookRepositoryMock) Find(id string) (book model.Book, apiError *errors.ApiError) {
	if r.TestError {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked test find error",
		}
	} else if r.TestNotFoundError {

		return model.Book{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "mocked test find not found error",
		}
	}

	book = model.Book{
		ID:             uint(25),
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		Available:      true,
	}

	return
}

func (r BookRepositoryMock) Create(book model.Book) (uint, *errors.ApiError) {
	_, err := r.BeforeCreate(book.CategoriesId)
	if r.TestCategoryNotFoundError {
		return 0, err
	} else if r.TestError {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   "mocked error",
		}
	}

	return 25, err
}

func (r BookRepositoryMock) Update(id string, upBook model.Book) (model.Book, *errors.ApiError) {
	err := r.BeforeUpdate(upBook.ID, upBook.CategoriesId)
	if r.TestNotFoundError {
		return model.Book{}, err
	} else if r.TestCategoryNotFoundError {
		return model.Book{}, err
	} else if r.TestError {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "mocked error",
		}
	}

	return model.Book{
		ID:             25,
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		Available:      true,
	}, err
}

func (r BookRepositoryMock) Delete(id string) (apiError *errors.ApiError) {
	r.BeforeDelete(1)
	if r.TestNotFoundError {
		return &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.DeleteFailMessage,
			Error:   "book not found",
		}
	}
	return nil
}

func (r BookRepositoryMock) BeforeCreate(categoriesId string) ([]uint, *errors.ApiError) {
	if r.TestCategoryNotFoundError {
		return []uint{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Error:   "category with ID: 5 not found",
			Message: errors.CreateFailMessage,
		}
	}

	r.AfterCreate(1, []uint{1})
	return []uint{}, nil
}

func (r BookRepositoryMock) AfterCreate(bookId uint, categoriesId []uint) {}

func (r BookRepositoryMock) BeforeUpdate(bookId uint, categoriesId string) *errors.ApiError {
	if r.TestNotFoundError {
		return &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.UpdateFailMessage,
			Error:   "book not found",
		}
	} else if r.TestCategoryNotFoundError {
		return &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.UpdateFailMessage,
			Error:   "category with ID: 5 not found",
		}
	}

	return nil
}

func (r BookRepositoryMock) BeforeDelete(bookId uint) {}
