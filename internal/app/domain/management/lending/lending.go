package lending

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

// Lending contains all Lending's table properties.
type Lending struct {
	ID        uint             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	BookID    uint             `json:"bookID" binding:"required" gorm:"unique"`
	Book      books.Book       `json:"book,omitempty" gorm:"->"`
	StudentID uint             `json:"studentID" binding:"required" gorm:"unique"`
	Student   students.Student `json:"student,omitempty" gorm:"->"`
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
