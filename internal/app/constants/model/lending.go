package model

import "time"

// Lending contains all Lending's table properties.
type Lending struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookID      uint      `json:"bookId" binding:"required"`
	StudentID   uint      `json:"studentId" binding:"required"`
	LendingDate time.Time `json:"lendingDate"`
}
