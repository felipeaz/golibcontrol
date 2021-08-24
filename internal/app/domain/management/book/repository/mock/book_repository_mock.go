package mock

import (
	"net/http"
	"os"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/model"
)

type BookRepositoryMock struct {
	TestError                     bool
	TestFindError                 bool
	TestCategoryNotFoundError     bool
	TestNotFoundError             bool
	TestBeforeUpdateNotFoundError bool
}

func (r BookRepositoryMock) Get() (books []model.Book, apiError *errors.ApiError) {
	if r.TestError {
		return nil, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	}

	books = []model.Book{
		model.Book{
			ID:             25,
			RegisterNumber: "123",
			Title:          "Mocked Book",
			Author:         "Mocked Author",
			Available:      true,
			CreatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
			UpdatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		},
	}

	return
}

func (r BookRepositoryMock) Find(id string) (book model.Book, apiError *errors.ApiError) {
	if r.TestFindError {
		return model.Book{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	} else if r.TestNotFoundError {

		return model.Book{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "book not found",
		}
	}

	book = model.Book{
		ID:             uint(25),
		RegisterNumber: "123",
		Title:          "Mocked Book",
		Author:         "Mocked Author",
		Available:      true,
		CreatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	return
}

func (r BookRepositoryMock) Create(book model.Book) (uint, *errors.ApiError) {
	_, apiError := r.BeforeCreate(book.CategoriesId)
	if r.TestCategoryNotFoundError {
		apiError.Message = errors.CreateFailMessage
		return 0, apiError
	} else if r.TestError {
		return 0, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   "mocked error",
		}
	}

	return 25, nil
}

func (r BookRepositoryMock) Update(id string, upBook model.Book) *errors.ApiError {
	_, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return apiError
	}

	err := r.BeforeUpdate(upBook.ID, upBook.CategoriesId)
	if r.TestCategoryNotFoundError || r.TestBeforeUpdateNotFoundError {
		return err
	} else if r.TestError {
		return &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "mocked error",
		}
	}

	return err
}

func (r BookRepositoryMock) Delete(id string) (apiError *errors.ApiError) {
	_, apiError = r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailMessage
		return
	}

	r.BeforeDelete(1)
	if r.TestError {
		return &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   "mocked error",
		}
	}
	return nil
}

func (r BookRepositoryMock) BeforeCreate(categoriesId string) ([]uint, *errors.ApiError) {
	if r.TestCategoryNotFoundError {
		return []uint{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
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
	if r.TestBeforeUpdateNotFoundError {
		return &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusNotFound,
			Message: errors.UpdateFailMessage,
			Error:   "book not found",
		}
	} else if r.TestCategoryNotFoundError {
		return &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusNotFound,
			Message: errors.UpdateFailMessage,
			Error:   "category with ID: 5 not found",
		}
	}

	return nil
}

func (r BookRepositoryMock) BeforeDelete(bookId uint) {}
