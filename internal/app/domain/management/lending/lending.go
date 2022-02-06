package lending

import (
	errorsx "errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
	"time"
)

var (
	BookUnavailableError    = errorsx.New("book unavailable")
	StudentHasActiveLending = errorsx.New("student already have an active lending")
)

// Lending contains all Lending's table properties.
type Lending struct {
	ID             uint             `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
	BookID         uint             `json:"bookID" binding:"required"`
	RegistryNumber int              `json:"registryNumber" binding:"required" gorm:"unique"`
	Book           books.Book       `json:"book,omitempty" gorm:"->"`
	StudentID      uint             `json:"studentID" binding:"required" gorm:"unique"`
	Student        students.Student `json:"student,omitempty" gorm:"->"`
}

func (l Lending) TableName() string {
	return "lending"
}

func (l *Lending) BeforeDelete(tx *gorm.DB) error {
	l.updateRegistry(tx, true)
	return nil
}

func (l *Lending) BeforeCreate(tx *gorm.DB) error {
	var respLending Lending
	resp := tx.Find(&respLending, "registry_number = ? OR student_id = ?", l.RegistryNumber, l.StudentID)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected != 0 {
		switch {
		case respLending.StudentID == l.StudentID:
			return StudentHasActiveLending
		case respLending.RegistryNumber == l.RegistryNumber:
			return BookUnavailableError
		}
	}

	resp = tx.Find(&registries.Registry{}, "registry_number = ?", l.RegistryNumber)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return BookUnavailableError
	}

	return nil
}

func (l *Lending) AfterCreate(tx *gorm.DB) error {
	l.updateRegistry(tx, false)
	return nil
}

func (l *Lending) updateRegistry(tx *gorm.DB, availability bool) {
	tx.Model(&registries.Registry{}).Where("registry_number = ?", l.RegistryNumber).
		Update("available", availability)
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
