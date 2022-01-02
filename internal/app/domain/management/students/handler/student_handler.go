package handler

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/gin-gonic/gin"
)

// StudentHandler handle the student router call.
type StudentHandler struct {
	Module students.Module
}

// NewStudentHandler Return an instance of this handler.
func NewStudentHandler(module students.Module) StudentHandler {
	return StudentHandler{
		Module: module,
	}
}

// Get returns all students.
func (h StudentHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Find return one student by ID.
func (h StudentHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create persist a student to the database.
func (h StudentHandler) Create(c *gin.Context) {
	authHeaderName := os.Getenv("AUTHORIZATION_TOKEN_NAME")
	accountHost := os.Getenv("API_GATEWAY_HOST")
	signinRoute := os.Getenv("SIGN_IN_URL")

	student, apiError := pkg.AssociateStudentInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(student, accountHost, signinRoute, authHeaderName, c.GetHeader(authHeaderName))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
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
