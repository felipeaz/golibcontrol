package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/repository/interface"
)

// StudentModule process the request recieved from handler.
type StudentModule struct {
	Repository _interface.StudentRepositoryInterface
}

func NewStudentModule(repo _interface.StudentRepositoryInterface) StudentModule {
	return StudentModule{
		Repository: repo,
	}
}

// Get returns all students.
func (m StudentModule) Get() ([]model.Student, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one student by ID.
func (m StudentModule) Find(id string) (model.Student, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a student to the database.
func (m StudentModule) Create(student model.Student) (string, *errors.ApiError) {
	return m.Repository.Create(student)
}

// Update update an existent student.
func (m StudentModule) Update(id string, upStudent model.Student) *errors.ApiError {
	return m.Repository.Update(id, upStudent)
}

// Delete delete an existent student.
func (m StudentModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
