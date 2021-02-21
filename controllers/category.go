package controllers

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/database/models"
	"github.com/gin-gonic/gin"
)

// GetCategories returns all categories.
func GetCategories(c *gin.Context) {
	var categories []models.Category

	models.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GetCategory return a single category.
func GetCategory(c *gin.Context) {
	var category models.Category

	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// CreateCategory insert category on DB.
func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	category := models.Category{Name: input.Name}
	models.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// UpdateCategory updates a category.
func UpdateCategory(c *gin.Context) {
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&category).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory updates a category.
func DeleteCategory(c *gin.Context) {
	var category models.Category

	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
