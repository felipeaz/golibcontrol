package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
)

// BookHandler handle the book router calls.
type BookHandler struct {
	Module module.BookModule
}

// NewBookHandler returns an instance of this handler.
func NewBookHandler(DB *gorm.DB) BookHandler {
	return BookHandler{
		Module: module.BookModule{
			Repository: repository.BookRepository{
				BookCategoryRepository: repository.BookCategoryRepository{
					DB: DB,
				},
				DB: DB,
			},
		},
	}
}

// Get returns all books.
func (h BookHandler) Get(c *gin.Context) {
	books, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Find return one book by ID.
func (h BookHandler) Find(c *gin.Context) {
	book, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Create creates a book.
func (h BookHandler) Create(c *gin.Context) {
	book, err := pkg.AssociateBookInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	id, apiError := h.Module.Create(book)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Update update an existent book.
func (h BookHandler) Update(c *gin.Context) {
	upBook, err := pkg.AssociateBookInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	book, apiError := h.Module.Update(c.Param("id"), upBook)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete delete an existent book by id.
func (h BookHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
