package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// StudentModule process the request recieved from handler.
type StudentModule struct {
	Repository repository.StudentRepository
}

// Get returns all students.
func (m StudentModule) Get() (students []model.Student, err error) {
	students, err = m.Repository.Get()
	return
}

// Find return one student by ID.
func (m StudentModule) Find(id int) (student model.Student, err error) {
	student, err = m.Repository.Find(id)
	return
}

// Create persist a student to the database.
func (m StudentModule) Create(student model.Student) (id uint, err error) {
	id, err = m.Repository.Create(student)
	return
}

// Update update an existent student.
func (m StudentModule) Update(id int, upStudent model.Student) (student model.Student, err error) {
	student, err = m.Repository.Update(id, upStudent)
	return
}

// Delete delete an existent student.
func (m StudentModule) Delete(id int) (err error) {
	err = m.Repository.Delete(id)
	return
}
