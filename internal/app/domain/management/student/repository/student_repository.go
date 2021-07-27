package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model/converter"
)

// StudentRepository is responsible of getting/saving information from DB.
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
	result, apiError := r.DB.FetchAll(&[]model.Student{})
	if apiError != nil {
		return nil, apiError
	}
	students, apiError := converter.ConvertToSliceStudentObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return students, apiError
}

// Find return one student from DB by ID.
func (r StudentRepository) Find(id string) (model.Student, *errors.ApiError) {
	result, apiError := r.DB.Fetch(&model.Student{}, id)
	if apiError != nil {
		return model.Student{}, apiError
	}

	student, apiError := converter.ConvertToStudentObj(result)
	if apiError != nil {
		return model.Student{}, apiError
	}

	return student, nil
}

// Create persist a student to the DB.
func (r StudentRepository) Create(student model.Student) (string, *errors.ApiError) {
	apiError := r.DB.Persist(&student)
	if apiError != nil {
		return "", apiError
	}

	return student.ID, nil
}

// Update update an existent student.
func (r StudentRepository) Update(id string, upStudent model.Student) *errors.ApiError {
	return r.DB.Refresh(&upStudent, id)
}

// Delete delete an existent student from DB.
func (r StudentRepository) Delete(id string) *errors.ApiError {
	apiError := r.DB.Remove(&model.Student{}, id)
	return apiError
}
