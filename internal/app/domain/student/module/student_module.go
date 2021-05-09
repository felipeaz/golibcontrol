package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/student/repository/interface"
)

// StudentModule process the request recieved from handler.
type StudentModule struct {
	Repository _interface.StudentRepositoryInterface
}

// Get returns all students.
func (m StudentModule) Get() (students []model.Student, apiError *errors.ApiError) {
	students, apiError = m.Repository.Get()
	return
}

// Find return one student by ID.
func (m StudentModule) Find(id string) (student model.Student, apiError *errors.ApiError) {
	student, apiError = m.Repository.Find(id)
	return
}

// Create persist a student to the database.
func (m StudentModule) Create(student model.Student) (id uint, apiError *errors.ApiError) {
	id, apiError = m.Repository.Create(student)
	return
}

// Update update an existent student.
func (m StudentModule) Update(id string, upStudent model.Student) (student model.Student, apiError *errors.ApiError) {
	student, apiError = m.Repository.Update(id, upStudent)
	return
}

// Delete delete an existent student.
func (m StudentModule) Delete(id string) (apiError *errors.ApiError) {
	apiError = m.Repository.Delete(id)
	return
}
