package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type myTime time.Time

var _ json.Unmarshaler = &myTime{}

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = myTime(t)
	return nil
}

// DB is an instance of gorm DB
var DB *gorm.DB

// Student contains all Student's table properties.
type Student struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	RegisterNumber string    `json:"registerNumber" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	Phone          string    `json:"phone" binding:"required"`
	Grade          string    `json:"grade" binding:"required"`
	Birthday       string    `json:"birthday" binding:"required"`
	CreatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
}

// Book contains all Book's table properties.
type Book struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title          string    `json:"title" binding:"required"`
	Author         string    `json:"author" binding:"required"`
	RegisterNumber string    `json:"registerNumber" binding:"required"`
	CategoryID     uint      `json:"categoryId" binding:"required"`
	CreatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
}

// Category contains all Category's table properties.
type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

// Lending contains all Lending's table properties.
type Lending struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookID      uint      `json:"bookId" binding:"required"`
	StudentID   uint      `json:"studentId" binding:"required"`
	LendingDate time.Time `json:"lendingDate"`
}
