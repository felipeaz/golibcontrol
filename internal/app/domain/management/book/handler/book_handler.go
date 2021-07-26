package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/module"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/module/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/repository"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/pkg"
	"github.com/gin-gonic/gin"
)

// BookHandler handle the book router calls.
type BookHandler struct {
	Module _interface.BookModuleInterface
}

// NewBookHandler returns an instance of this handler.
func NewBookHandler(dbService *service.MySQLService) BookHandler {
	return BookHandler{
		Module: module.NewBookModule(
			repository.NewBookRepository(dbService,
				repository.NewBookCategoryRepository(dbService)),
		),
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
	book, apiError := pkg.AssociateBookInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
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
