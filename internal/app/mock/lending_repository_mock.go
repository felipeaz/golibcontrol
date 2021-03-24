package mock

import (
	"net/http"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type LendingRepositoryMock struct {
	TestError                   bool
	TestNotFoundError           bool
	TestBookNotFoundError       bool
	TestStudentNotFoundError    bool
	TestBookAlreadyLentError    bool
	TestStudentAlreadyLentError bool
}

// Get returns all lendings.
func (r LendingRepositoryMock) Get() (lendings []model.Lending, apiError *errors.ApiError) {
	if r.TestError {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	}

	lendings = []model.Lending{
		model.Lending{
			ID:          25,
			BookID:      5,
			StudentID:   10,
			LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		},
	}

	return lendings, nil
}

// Find return one lending from DB by ID.
func (r LendingRepositoryMock) Find(id string) (lending model.Lending, apiError *errors.ApiError) {
	if r.TestError {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	}

	if r.TestNotFoundError {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "lending not found",
		}
	}

	lending = model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	return lending, nil
}

// Create persist a lending to the DB.
func (r LendingRepositoryMock) Create(lending model.Lending) (uint, *errors.ApiError) {
	if apiError := r.BeforeCreateAndUpdate(lending.StudentID, lending.BookID); apiError != nil {
		apiError.Message = errors.CreateFailMessage
		return 0, apiError
	}

	if apiError := r.BeforeCreate(lending.StudentID, lending.BookID); apiError != nil {
		apiError.Message = errors.CreateFailMessage
		return 0, apiError
	}

	if r.TestError {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   "mocked error",
		}
	}

	return 25, nil
}

// Update update an existent lending.
func (r LendingRepositoryMock) Update(id string, upLending model.Lending) (model.Lending, *errors.ApiError) {
	if r.TestNotFoundError {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.UpdateFailMessage,
			Error:   "lending not found",
		}
	}

	apiError := r.BeforeCreateAndUpdate(upLending.StudentID, upLending.BookID)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return model.Lending{}, apiError
	}

	if r.TestError {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "mocked error",
		}
	}

	return upLending, nil
}

// Delete delete an existent lending from DB.
func (r LendingRepositoryMock) Delete(id string) (apiError *errors.ApiError) {
	if r.TestNotFoundError {
		return &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.DeleteFailMessage,
			Error:   "lending not found",
		}
	}

	if r.TestError {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   "mocked error",
		}
	}

	return nil
}

// beforeCreateAndUpdate validate if the student or book exists before create the lending.
func (r LendingRepositoryMock) BeforeCreateAndUpdate(studentId, bookId uint) *errors.ApiError {
	if r.TestStudentNotFoundError {
		if !r.TestError {
			return &errors.ApiError{
				Status: http.StatusNotFound,
				Error:  "student not found",
			}
		}

		if r.TestError {
			return &errors.ApiError{
				Status: http.StatusInternalServerError,
				Error:  "mocked error",
			}
		}
	}

	if r.TestBookNotFoundError {
		if !r.TestError {
			return &errors.ApiError{
				Status: http.StatusNotFound,
				Error:  "book not found",
			}
		}

		return &errors.ApiError{
			Status: http.StatusInternalServerError,
			Error:  "mocked error",
		}
	}

	return nil
}

// beforeCreate validate if the book is already lent.
func (r LendingRepositoryMock) BeforeCreate(studentId, bookId uint) *errors.ApiError {
	if r.TestBookAlreadyLentError {
		return &errors.ApiError{
			Status: http.StatusInternalServerError,
			Error:  "book is already lent",
		}
	}

	if r.TestStudentAlreadyLentError {
		return &errors.ApiError{
			Status: http.StatusInternalServerError,
			Error:  "student has already lent a book",
		}
	}

	return nil
}
