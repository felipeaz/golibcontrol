package handler

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/module"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/module/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/repository"
	"github.com/gin-gonic/gin"
)

// StudentHandler handle the student router call.
type StudentHandler struct {
	Module _interface.StudentModuleInterface
}

// NewStudentHandler Return an instance of this handler.
func NewStudentHandler(dbService *service.MySQLService) StudentHandler {
	return StudentHandler{
		Module: module.NewStudentModule(repository.NewStudentRepository(dbService)),
	}
}

// Get returns all students.
func (h StudentHandler) Get(c *gin.Context) {
	students, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": students})
}

// Find return one student by ID.
func (h StudentHandler) Find(c *gin.Context) {
	student, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// Create persist a student to the database.
func (h StudentHandler) Create(c *gin.Context) {
	student, apiError := pkg.AssociateStudentInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	accountHost := os.Getenv("API_GATEWAY_HOST") + os.Getenv("SIGN_IN_URL")
	id, apiError := h.Module.Create(student, accountHost)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Update update an existent student.
func (h StudentHandler) Update(c *gin.Context) {
	upStudent, apiError := pkg.AssociateStudentInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upStudent)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}

// Delete delete an existent student.
func (h StudentHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}
