package repository

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// StudentRepository is responsible of getting/saving information from DB.
type StudentRepository struct {
	DB *gorm.DB
}

// Get returns all students.
func (r StudentRepository) Get() (students []model.Student, apiError *errors.ApiError) {
	result := r.DB.Find(&students)
	if err := result.Error; err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	return
}

// Find return one student from DB by ID.
func (r StudentRepository) Find(id string) (student model.Student, apiError *errors.ApiError) {
	result := r.DB.Model(model.Student{}).Where("id = ?", id).First(&student)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return model.Student{}, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}

		return model.Student{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "student not found",
		}
	}

	return
}

// Create persist a student to the DB.
func (r StudentRepository) Create(student model.Student) (uint, *errors.ApiError) {
	result := r.DB.Create(&student)
	if err := result.Error; err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return student.ID, nil
}

// Update update an existent student.
func (r StudentRepository) Update(id string, upStudent model.Student) (model.Student, *errors.ApiError) {
	student, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return model.Student{}, apiError
	}

	result := r.DB.Model(&student).Updates(upStudent)
	if err := result.Error; err != nil {
		return model.Student{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}

	return student, nil
}

// Delete delete an existent student from DB.
func (r StudentRepository) Delete(id string) (apiError *errors.ApiError) {
	student, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailMessage
		return
	}

	err := r.DB.Delete(&student).Error
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}

	return
}
