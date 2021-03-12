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
