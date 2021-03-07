package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
				DB: DB,
			},
		},
	}
}

// GetBooks returns all books.
func (h BookHandler) GetBooks(c *gin.Context) {
	books, err := h.Module.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GetBook return one book by ID.
func (h BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a number"})
		return
	}

	book, err := h.Module.GetBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreatBook creates a book.
func (h BookHandler) CreateBook(c *gin.Context) {
	book, err := associateInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Module.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// UpdateBook update an existent book.
func (h BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a number"})
		return
	}

	upBook, err := associateInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Module.UpdateBook(id, upBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook delete an existent book by id.
func (h BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a number"})
		return
	}

	err = h.Module.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

// associateInput is responsible of associate the params to the model.
func associateInput(c *gin.Context) (book model.Book, err error) {
	err = c.ShouldBindJSON(&book)
	return
}
