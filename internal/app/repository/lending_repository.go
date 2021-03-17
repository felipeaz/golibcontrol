package repository

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// LendingRepository is responsible of getting/saving information from DB.
type LendingRepository struct {
	DB                *gorm.DB
	StudentRepository StudentRepository
	BookRepository    BookRepository
}

// Get returns all lendings.
func (r LendingRepository) Get() (lendings []model.Lending, apiError *errors.ApiError) {
	result := r.DB.Find(&lendings)
	if err := result.Error; err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.NotFoundMessage,
			Error:   err.Error(),
		}
	}

	return
}

// Find return one lending from DB by ID.
func (r LendingRepository) Find(id string) (lending model.Lending, apiError *errors.ApiError) {
	result := r.DB.Model(model.Lending{}).Where("id = ?", id).First(&lending)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return model.Lending{}, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.NotFoundMessage,
				Error:   err.Error(),
			}
		}

		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.NotFoundMessage,
			Error:   "lending not found",
		}
	}

	return
}

// Create persist a lending to the DB.
func (r LendingRepository) Create(lending model.Lending) (uint, *errors.ApiError) {
	err := r.beforeCreateAndUpdate(lending.StudentID, lending.BookID)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailedMessage,
			Error:   err.Error(),
		}
	}

	err = r.beforeCreate(lending.StudentID, lending.BookID)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailedMessage,
			Error:   err.Error(),
		}
	}

	result := r.DB.Create(&lending)
	if err = result.Error; err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailedMessage,
			Error:   err.Error(),
		}
	}

	return lending.ID, nil
}

// Update update an existent lending.
func (r LendingRepository) Update(id string, upLending model.Lending) (model.Lending, *errors.ApiError) {
	lending, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailedMessage
		return model.Lending{}, apiError
	}

	err := r.beforeCreateAndUpdate(upLending.StudentID, upLending.BookID)
	if err != nil {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailedMessage,
			Error:   err.Error(),
		}
	}

	result := r.DB.Model(&lending).Updates(upLending)
	if err = result.Error; err != nil {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailedMessage,
			Error:   err.Error(),
		}
	}

	return lending, nil
}

// Delete delete an existent lending from DB.
func (r LendingRepository) Delete(id string) (apiError *errors.ApiError) {
	lending, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailedMessage
		return
	}

	err := r.DB.Delete(&lending).Error
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailedMessage,
			Error:   err.Error(),
		}
	}

	return
}

// beforeCreateAndUpdate validate if the student or book exists before create the lending.
func (r LendingRepository) beforeCreateAndUpdate(studentId, bookId uint) error {
	var student model.Student
	result := r.DB.First(&student, studentId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("student not found")
		}

		return err
	}

	var book model.Book
	result = r.DB.First(&book, bookId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("book not found")
		}

		return err
	}

	return nil
}

// beforeCreate validate if the book is already lent.
func (r LendingRepository) beforeCreate(studentId, bookId uint) error {
	var lending model.Lending
	result := r.DB.Where("book_id = ?", bookId).First(&lending)
	if result.RowsAffected > 0 {
		return fmt.Errorf("book is already lent")
	}

	result = r.DB.Where("student_id = ?", studentId).First(&lending)
	if result.RowsAffected > 0 {
		return fmt.Errorf("student has already lent a book")
	}

	return nil
}
