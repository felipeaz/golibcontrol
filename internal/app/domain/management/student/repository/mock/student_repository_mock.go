package mock

import (
	"net/http"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
)

type StudentRepositoryMock struct {
	TestError         bool
	TestNotFoundError bool
	TestUpdateError   bool
	TestDeleteError   bool
}

// Get returns all students.
func (r StudentRepositoryMock) Get() (students []model.Student, apiError *errors.ApiError) {
	if r.TestError {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	}

	students = []model.Student{
		model.Student{
			ID:             25,
			RegisterNumber: "2500651",
			Name:           "Felipe de Azevedo Silva",
			Email:          "felipe9_azevedo@hotmail.com",
			Phone:          "(00)00000-0000",
			Grade:          "7th",
			Birthday:       "31/12/1997",
			CreatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
			UpdatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		},
	}

	return students, nil
}

// Find return one student from DB by ID.
func (r StudentRepositoryMock) Find(id string) (student model.Student, apiError *errors.ApiError) {
	if r.TestError {
		return model.Student{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	}

	if r.TestNotFoundError {
		return model.Student{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "student not found",
		}
	}

	student = model.Student{
		ID:             25,
		RegisterNumber: "2500651",
		Name:           "Felipe de Azevedo Silva",
		Email:          "felipe9_azevedo@hotmail.com",
		Phone:          "(00)00000-0000",
		Grade:          "7th",
		Birthday:       "31/12/1997",
		CreatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt:      time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	return student, nil
}

// Create persist a student to the DB.
func (r StudentRepositoryMock) Create(student model.Student) (uint, *errors.ApiError) {
	if r.TestError {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   "mocked error",
		}
	}

	return 25, nil
}

// Update update an existent student.
func (r StudentRepositoryMock) Update(id string, upStudent model.Student) *errors.ApiError {
	_, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return apiError
	}

	if r.TestUpdateError {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "mocked error",
		}
	}

	return nil
}

// Delete delete an existent student from DB.
func (r StudentRepositoryMock) Delete(id string) (apiError *errors.ApiError) {
	_, apiError = r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailMessage
		return
	}

	if r.TestDeleteError {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   "mocked error",
		}
	}

	return nil
}
