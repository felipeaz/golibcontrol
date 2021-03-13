package pkg

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

// AssociateBookInput is responsible of associate the params to the book model.
func AssociateBookInput(c *gin.Context) (book model.Book, err error) {
	err = c.ShouldBindJSON(&book)
	return
}

// AssociateCategoryInput is responsible of associating the params to the category model.
func AssociateCategoryInput(c *gin.Context) (category model.Category, err error) {
	err = c.ShouldBindJSON(&category)
	return
}

// AssociateBookCategoryInput is responsible of associating the params to the bookCategory model.
func AssociateBookCategoryInput(c *gin.Context) (bookCategory model.BookCategory, err error) {
	err = c.ShouldBindJSON(&bookCategory)
	return
}

// AssociateStudentInput is responsible of associating the params to the student model.
func AssociateStudentInput(c *gin.Context) (student model.Student, err error) {
	err = c.ShouldBindJSON(&student)
	return
}

// AssociateLendingInput is responsible of associating the params to the lending model.
func AssociateLendingInput(c *gin.Context) (lending model.Lending, err error) {
	err = c.ShouldBindJSON(&lending)
	return
}
