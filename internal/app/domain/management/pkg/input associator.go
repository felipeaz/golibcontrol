package pkg

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/model"
	categoryModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/model"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	studentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
	"github.com/gin-gonic/gin"
)

// AssociateBookInput is responsible of associate the params to the book model.
func AssociateBookInput(c *gin.Context) (book bookModel.Book, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&book)
	if err != nil {
		return bookModel.Book{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateCategoryInput is responsible of associating the params to the category model.
func AssociateCategoryInput(c *gin.Context) (category categoryModel.Category, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return categoryModel.Category{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateStudentInput is responsible of associating the params to the student model.
func AssociateStudentInput(c *gin.Context) (student studentModel.Student, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&student)
	if err != nil {
		return studentModel.Student{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateLendingInput is responsible of associating the params to the lending model.
func AssociateLendingInput(c *gin.Context) (lending lendingModel.Lending, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&lending)
	if err != nil {
		return lendingModel.Lending{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
