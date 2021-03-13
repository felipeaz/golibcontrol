package handler

import (
	"net/http"
	"strconv"

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
			BookRepository: repository.BookRepository{
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
	books, err := h.Module.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Find return one book by ID.
func (h BookHandler) Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Module.Find(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Create creates a book.
func (h BookHandler) Create(c *gin.Context) {
	book, err := pkg.AssociateBookInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Module.Create(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Update update an existent book.
func (h BookHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upBook, err := pkg.AssociateBookInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Module.Update(id, upBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete delete an existent book by id.
func (h BookHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Module.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
