package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/app/management/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookHandler handle the book router calls.
type BookHandler struct {
	Module books.Module
}

// NewBookHandler returns an instance of this handler.
func NewBookHandler(module books.Module) BookHandler {
	return BookHandler{
		Module: module,
	}
}

// Get returns all books.
func (h BookHandler) Get(c *gin.Context) {
	var params books.Filter
	if err := c.Bind(&params); err == nil {
		resp, apiError := h.Module.GetByFilter(params)
		if apiError != nil {
			c.JSON(apiError.Status, apiError)
			return
		}

		c.JSON(http.StatusOK, resp)
		return
	}

	resp, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Find return one book by ID.
func (h BookHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
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
