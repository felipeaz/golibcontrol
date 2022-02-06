package registries

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

type Registry struct {
	RegistryNumber int       `json:"registryNumber" gorm:"primaryKey;not null;unique"`
	BookID         uint      `json:"bookId" gorm:"not null"`
	Available      bool      `json:"available" gorm:"default:false"`
	Reserved       bool      `json:"reserved" gorm:"default:false"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (r Registry) TableName() string {
	return "book_registry_numbers"
}

type Filter struct {
	BookID         string `form:"bookId,omitempty" column:"book_registry_numbers.book_id" array:"false"`
	RegistryNumber string `form:"registryNumber,omitempty" column:"book_registry_numbers.registry_number" array:"false"`
	Available      string `form:"available,omitempty" column:"book_registry_numbers.available" array:"false"`
	Reserved       string `form:"reserved,omitempty" column:"book_registry_numbers.reserved" array:"false"`
}

func (f Filter) GetFieldNames() []string {
	return []string{"BookID", "RegistryNumber", "Available"}
}

type Module interface {
	Get() ([]Registry, *errors.ApiError)
	GetByFilter(filter Filter) ([]Registry, *errors.ApiError)
	Find(id string) (Registry, *errors.ApiError)
	Create(category Registry) (*Registry, *errors.ApiError)
	Update(id string, upRegistry Registry) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() ([]Registry, *errors.ApiError)
	GetByFilter(filter Filter) ([]Registry, *errors.ApiError)
	Find(id string) (Registry, *errors.ApiError)
	Create(category Registry) (*Registry, *errors.ApiError)
	Update(id string, upRegistry Registry) *errors.ApiError
	Delete(id string) *errors.ApiError
}
