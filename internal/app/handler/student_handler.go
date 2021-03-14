package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// StudentHandler handle the student router call.
type StudentHandler struct {
	Module module.StudentModule
}

// NewStudentHandler Return an instance of this handler.
func NewStudentHandler(db *gorm.DB) StudentHandler {
	return StudentHandler{
		Module: module.StudentModule{
			Repository: repository.StudentRepository{
				DB: db,
			},
		},
	}
}

// Get returns all students.
func (h StudentHandler) Get(c *gin.Context) {

}

// Find return one student by ID.
func (h StudentHandler) Find(c *gin.Context) {

}

// Create creates a student.
func (h StudentHandler) Create(c *gin.Context) {

}

// Update update an existent student.
func (h StudentHandler) Update(c *gin.Context) {

}

// Delete delete an existent student by id.
func (h StudentHandler) Delete(c *gin.Context) {

}
