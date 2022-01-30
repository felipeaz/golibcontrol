package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AssociateBookInput is responsible for associate the params to the book model.
func AssociateBookInput(c *gin.Context) (book books.Book, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&book)
	if err != nil {
		return books.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateCategoryInput is responsible for associating the params to the category model.
func AssociateCategoryInput(c *gin.Context) (category categories.Category, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return categories.Category{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateRegistryInput is responsible for associating the params to the registry model.
func AssociateRegistryInput(c *gin.Context) (registry registries.Registry, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&registry)
	if err != nil {
		return registries.Registry{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateStudentInput is responsible for associating the params to the student model.
func AssociateStudentInput(c *gin.Context) (student students.Student, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&student)
	if err != nil {
		return students.Student{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateLendingInput is responsible for associating the params to the lending model.
func AssociateLendingInput(c *gin.Context) (lend lending.Lending, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&lend)
	if err != nil {
		return lending.Lending{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
