package controllers

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/database/models"
	"github.com/gin-gonic/gin"
)

func validateCategory(c *gin.Context, input models.Book) (bool, string) {
	var category models.Category
	if err := models.DB.Where("id = ?", input.CategoryID).First(&category).Error; err != nil {
		return false, "Category Not Found"
	}

	return true, ""
}

// GetBooks return all books.
func GetBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GetBook return an specific book.
func GetBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook persist book into the database.
func CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if status, msg := validateCategory(c, input); status == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	book := models.Book{
		Title:          input.Title,
		Author:         input.Author,
		RegisterNumber: input.RegisterNumber,
		Available:      input.Available,
		CategoryID:     input.CategoryID,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook update a book.
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if status, msg := validateCategory(c, input); status == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	if err := models.DB.Model(&book).Updates(map[string]interface{}{
		"Title":          input.Title,
		"Author":         input.Author,
		"RegisterNumber": input.RegisterNumber,
		"Available":      input.Available,
		"CategoryID":     input.CategoryID,
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook delete a book.
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
