package repository

import (
	"gorm.io/gorm"
)

// StudentRepository is responsible of getting/saving information from DB.
type StudentRepository struct {
	DB *gorm.DB
}
