package controllers

import (
	"net/http"
	"time"

	"github.com/FelipeAz/golibcontrol/database/models"
	"github.com/gin-gonic/gin"
)

// GetStudents return all students.
func GetStudents(c *gin.Context) {
	var student []models.Student

	models.DB.Find(&student)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// GetStudent return one studant.
func GetStudent(c *gin.Context) {
	var student models.Student
	if err := models.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// CreateStudent persist a student into database.
func CreateStudent(c *gin.Context) {
	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	birthStr := input.Birthday
	dateFormated, err := time.Parse("02/01/2006", birthStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	student := models.Student{
		RegisterNumber: input.RegisterNumber,
		Name:           input.Name,
		Email:          input.Email,
		Phone:          input.Phone,
		Birthday:       dateFormated.Format("02/01/2006"),
		Grade:          input.Grade,
	}
	models.DB.Create(&student)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// UpdateStudent update a specific student.
func UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := models.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&student).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// DeleteStudent deletes a student.
func DeleteStudent(c *gin.Context) {
	var student models.Student
	if err := models.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&student)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
