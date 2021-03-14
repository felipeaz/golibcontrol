package repository

import (
	"gorm.io/gorm"
)

// LendingRepository is responsible of getting/saving information from DB.
type LendingRepository struct {
	DB *gorm.DB
}
