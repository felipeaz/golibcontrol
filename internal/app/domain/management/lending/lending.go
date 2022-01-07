package lending

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"gorm.io/gorm"
)

// Lending contains all Lending's table properties.
type Lending struct {
	gorm.Model
	BookID    uint `json:"bookId" binding:"required" gorm:"unique"`
	StudentID uint `json:"studentId" binding:"required" gorm:"unique"`
}

func (l Lending) TableName() string {
	return "lending"
}

type Module interface {
	Get() ([]Lending, *errors.ApiError)
	Find(id string) (Lending, *errors.ApiError)
	Create(lending Lending) (*Lending, *errors.ApiError)
	Update(id string, upLending Lending) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() (lendings []Lending, apiError *errors.ApiError)
	Find(id string) (lending Lending, apiError *errors.ApiError)
	Create(lending Lending) (*Lending, *errors.ApiError)
	Update(id string, upLending Lending) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
