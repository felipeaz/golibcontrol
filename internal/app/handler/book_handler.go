package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BookHandler handle the book router calls.
type BookHandler struct {
	module module.BookModule
}

// GetBooks returns all books.
func (h BookHandler) GetBooks(c *gin.Context) {
	books, err := h.module.GetBooks()
	responseHandler(books, err, c)
}

// GetBook return one book by ID.
func (h BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a number"})
		return
	}

	book, err := h.module.GetBook(id)
	responseHandler(book, err, c)
}

// CreatBook creates a book.
func (h BookHandler) CreateBook(c *gin.Context) {
	book, err := associateInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.module.CreateBook(book)
	responseHandler(id, err, c)
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

	book, err := h.module.UpdateBook(id, upBook)
	responseHandler(book, err, c)
}

// DeleteBook delete an existent book by id.
func (h BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a number"})
		return
	}

	status, err := h.module.DeleteBook(id)
	responseHandler(status, err, c)
}

// responseHandler deals with json responses validating the error.
func responseHandler(data interface{}, err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// associateInput is responsible of associate the params to the model.
func associateInput(c *gin.Context) (book model.Book, err error) {
	err = c.ShouldBindJSON(&book)
	return
}
