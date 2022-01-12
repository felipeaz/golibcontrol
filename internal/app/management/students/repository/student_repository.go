package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/FelipeAz/golibcontrol/internal/app/management/students/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
func (r StudentRepository) Get() ([]students.Student, *errors.ApiError) {
	result, err := r.DB.Find(nil, &[]students.Student{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceStudentObj(result)
}

// Find return one student from DB by ID.
func (r StudentRepository) Find(id string) (students.Student, *errors.ApiError) {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &students.Student{})
	if err != nil {
		return students.Student{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToStudentObj(result)
}

// Create persist a student to the DB.
func (r StudentRepository) Create(student students.Student) (*students.Student, *errors.ApiError) {
	err := r.DB.Persist(&student)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &student, nil
}

// Update update an existent student.
func (r StudentRepository) Update(id string, upStudent students.Student) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upStudent)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &students.Student{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
