package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model/converter"
)

// StudentRepository is responsible for getting/saving information from DB.
type StudentRepository struct {
	DB database.GORMServiceInterface
}

func NewStudentRepository(db database.GORMServiceInterface) StudentRepository {
	return StudentRepository{
		DB: db,
	}
}

// Get returns all students.
func (r StudentRepository) Get() ([]model.Student, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]model.Student{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	students, apiError := converter.ConvertToSliceStudentObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return students, apiError
}

// Find return one student from DB by ID.
func (r StudentRepository) Find(id string) (model.Student, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Student{}, id)
	if err != nil {
		return model.Student{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	student, apiError := converter.ConvertToStudentObj(result)
	if apiError != nil {
		return model.Student{}, apiError
	}

	return student, nil
}

// Create persist a student to the DB.
func (r StudentRepository) Create(student model.Student) (string, *errors.ApiError) {
	err := r.DB.Persist(&student)
	if err != nil {
		return "", &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return student.ID, nil
}

// Update update an existent student.
func (r StudentRepository) Update(id string, upStudent model.Student) *errors.ApiError {
	err := r.DB.Refresh(&upStudent, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// Delete delete an existent student from DB.
func (r StudentRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Student{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
