package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// LendingRepository is responsible of getting/saving information from DB.
type LendingRepository struct {
	DB                *gorm.DB
	StudentRepository StudentRepository
	BookRepository    BookRepository
}

// Get returns all lendings.
func (r LendingRepository) Get() (lendings []model.Lending, err error) {
	result := r.DB.Find(&lendings)
	if err = result.Error; err != nil {
		return nil, err
	}

	return
}

// Find return one lending from DB by ID.
func (r LendingRepository) Find(id int) (lending model.Lending, err error) {
	result := r.DB.Model(model.Lending{}).Where("id = ?", id).First(&lending)
	if err = result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Lending{}, fmt.Errorf("lending not found")
		}

		return model.Lending{}, err
	}

	return
}

// Create persist a lending to the DB.
func (r LendingRepository) Create(lending model.Lending) (uint, error) {
	err := r.beforeCreateAndUpdate(lending.StudentID, lending.BookID)
	if err != nil {
		return 0, err
	}

	err = r.beforeCreate(lending.BookID)
	if err != nil {
		return 0, err
	}

	result := r.DB.Create(&lending)
	if err = result.Error; err != nil {
		return 0, err
	}

	return lending.ID, nil
}

// Update update an existent lending.
func (r LendingRepository) Update(id int, upLending model.Lending) (model.Lending, error) {
	lending, err := r.Find(id)
	if err != nil {
		return model.Lending{}, err
	}

	err = r.beforeCreateAndUpdate(upLending.StudentID, upLending.BookID)
	if err != nil {
		return model.Lending{}, err
	}

	result := r.DB.Model(&lending).Updates(upLending)
	if err = result.Error; err != nil {
		return model.Lending{}, err
	}

	return lending, nil
}

// Delete delete an existent lending from DB.
func (r LendingRepository) Delete(id int) (err error) {
	lending, err := r.Find(id)
	if err != nil {
		return
	}

	err = r.DB.Delete(&lending).Error
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
func (r LendingRepository) beforeCreate(bookId uint) error {
	var lending model.Lending
	result := r.DB.Where("book_id = ?", bookId).First(&lending)
	if result.RowsAffected > 0 {
		return fmt.Errorf("book is already lent")
	}

	return nil
}
