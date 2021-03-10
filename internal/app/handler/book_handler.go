package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"

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
func (h BookHandler) Get(c *gin.Context) {
	books, err := h.Module.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GetBook return one book by ID.
func (h BookHandler) Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a number"})
		return
	}

	book, err := h.Module.Find(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreatBook creates a book.
func (h BookHandler) Create(c *gin.Context) {
	book, err := associateInput(c)
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

// UpdateBook update an existent book.
func (h BookHandler) Update(c *gin.Context) {
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

	book, err := h.Module.Update(id, upBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook delete an existent book by id.
func (h BookHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should be a number"})
		return
	}

	err = h.Module.Delete(id)
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
