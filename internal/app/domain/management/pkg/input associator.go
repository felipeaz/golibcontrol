package pkg

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
	categoryModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories/model"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	studentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/students/model"
	"github.com/gin-gonic/gin"
)

// AssociateBookInput is responsible for associate the params to the book model.
func AssociateBookInput(c *gin.Context) (book bookModel.Book, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&book)
	if err != nil {
		return bookModel.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateCategoryInput is responsible for associating the params to the category model.
func AssociateCategoryInput(c *gin.Context) (category categoryModel.Category, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return categoryModel.Category{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateStudentInput is responsible for associating the params to the student model.
func AssociateStudentInput(c *gin.Context) (student studentModel.Student, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&student)
	if err != nil {
		return studentModel.Student{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateLendingInput is responsible for associating the params to the lending model.
func AssociateLendingInput(c *gin.Context) (lending lendingModel.Lending, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&lending)
	if err != nil {
		return lendingModel.Lending{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
