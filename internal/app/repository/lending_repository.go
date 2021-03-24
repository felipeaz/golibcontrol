package repository

import (
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
			Message: errors.FailMessage,
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
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}

		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "lending not found",
		}
	}

	return
}

// Create persist a lending to the DB.
func (r LendingRepository) Create(lending model.Lending) (uint, *errors.ApiError) {
	if apiError := r.BeforeCreateAndUpdate(lending.StudentID, lending.BookID); apiError != nil {
		apiError.Message = errors.CreateFailMessage
		return 0, apiError
	}

	if apiError := r.BeforeCreate(lending.StudentID, lending.BookID); apiError != nil {
		apiError.Message = errors.CreateFailMessage
		return 0, apiError
	}

	result := r.DB.Create(&lending)
	if err := result.Error; err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return lending.ID, nil
}

// Update update an existent lending.
func (r LendingRepository) Update(id string, upLending model.Lending) (model.Lending, *errors.ApiError) {
	lending, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return model.Lending{}, apiError
	}

	apiError = r.BeforeCreateAndUpdate(upLending.StudentID, upLending.BookID)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return model.Lending{}, apiError
	}

	result := r.DB.Model(&lending).Updates(upLending)
	if err := result.Error; err != nil {
		return model.Lending{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}

	return lending, nil
}

// Delete delete an existent lending from DB.
func (r LendingRepository) Delete(id string) (apiError *errors.ApiError) {
	lending, apiError := r.Find(id)
	if apiError != nil {
		apiError.Message = errors.DeleteFailMessage
		return
	}

	err := r.DB.Delete(&lending).Error
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}

	return
}

// beforeCreateAndUpdate validate if the student or book exists before create the lending.
func (r LendingRepository) BeforeCreateAndUpdate(studentId, bookId uint) *errors.ApiError {
	var student model.Student
	result := r.DB.First(&student, studentId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &errors.ApiError{
				Status:  http.StatusNotFound,
				Message: errors.UpdateFailMessage,
				Error:   "student not found",
			}
		}

		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}

	var book model.Book
	result = r.DB.First(&book, bookId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &errors.ApiError{
				Status:  http.StatusNotFound,
				Message: errors.UpdateFailMessage,
				Error:   "book not found",
			}
		}

		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}

	return nil
}

// beforeCreate validate if the book is already lent.
func (r LendingRepository) BeforeCreate(studentId, bookId uint) *errors.ApiError {
	var lending model.Lending
	result := r.DB.Where("book_id = ?", bookId).First(&lending)
	if result.RowsAffected > 0 {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "book is already lent",
		}
	}

	result = r.DB.Where("student_id = ?", studentId).First(&lending)
	if result.RowsAffected > 0 {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "student has already lent a book",
		}
	}

	return nil
}
