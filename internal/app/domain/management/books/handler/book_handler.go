package handler

import (
	domain "github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/pkg"
	"github.com/gin-gonic/gin"
)

// BookHandler handle the book router calls.
type BookHandler struct {
	Module domain.Module
}

// NewBookHandler returns an instance of this handler.
func NewBookHandler(module domain.Module) BookHandler {
	return BookHandler{
		Module: module,
	}
}

// Get returns all domain.
func (h BookHandler) Get(c *gin.Context) {
	var params domain.Filter
	if err := c.Bind(&params); err == nil {
		books, apiError := h.Module.GetByFilter(params)
		if apiError != nil {
			c.JSON(apiError.Status, apiError)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": books})
		return
	}

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
	book, apiError := pkg.AssociateBookInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(book)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update update an existent book.
func (h BookHandler) Update(c *gin.Context) {
	upBook, apiError := pkg.AssociateBookInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upBook)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}

// Delete delete an existent book by id.
func (h BookHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}
