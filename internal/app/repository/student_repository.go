package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// StudentRepository is responsible of getting/saving information from DB.
type StudentRepository struct {
	DB *gorm.DB
}

// Get returns all students.
func (r StudentRepository) Get() (students []model.Student, err error) {
	result := r.DB.Find(&students)
	if err = result.Error; err != nil {
		return nil, err
	}

	return
}

// Find return one student from DB by ID.
func (r StudentRepository) Find(id int) (student model.Student, err error) {
	result := r.DB.Model(model.Student{}).Where("id = ?", id).First(&student)
	if err = result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Student{}, fmt.Errorf("student not found")
		}

		return model.Student{}, err
	}

	return
}

// Create persist a student to the DB.
func (r StudentRepository) Create(student model.Student) (uint, error) {
	result := r.DB.Create(&student)
	if err := result.Error; err != nil {
		return 0, err
	}

	return student.ID, nil
}

// Update update an existent student.
func (r StudentRepository) Update(id int, upStudent model.Student) (model.Student, error) {
	student, err := r.Find(id)
	if err != nil {
		return model.Student{}, err
	}

	result := r.DB.Model(&student).Updates(upStudent)
	if err := result.Error; err != nil {
		return model.Student{}, err
	}

	return student, nil
}

// Delete delete an existent student from DB.
func (r StudentRepository) Delete(id int) (err error) {
	student, err := r.Find(id)
	if err != nil {
		return
	}

	err = r.DB.Delete(&student).Error
	return
}
