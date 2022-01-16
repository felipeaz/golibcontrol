package lending

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

// Lending contains all Lending's table properties.
type Lending struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	BookID    uint      `json:"bookId" binding:"required" gorm:"unique"`
	StudentID uint      `json:"studentId" binding:"required" gorm:"unique"`
}

func (l Lending) TableName() string {
	return "lending"
}

type Filter struct {
	StudentID string `form:"studentId,omitempty" column:"lending.student_id"`
	BookID    string `form:"bookId,omitempty" column:"lending.book_id"`
}

func (f Filter) GetFieldNames() []string {
	return []string{"BookID", "StudentID"}
}

type Module interface {
	Get() ([]Lending, *errors.ApiError)
	GetByFilter(filter Filter) ([]Lending, *errors.ApiError)
	Find(id string) (Lending, *errors.ApiError)
	Create(lending Lending) (*Lending, *errors.ApiError)
	Update(id string, upLending Lending) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() ([]Lending, *errors.ApiError)
	GetByFilter(filter Filter) ([]Lending, *errors.ApiError)
	Find(id string) (Lending, *errors.ApiError)
	Create(lending Lending) (*Lending, *errors.ApiError)
	Update(id string, upLending Lending) *errors.ApiError
	Delete(id string) *errors.ApiError
}
