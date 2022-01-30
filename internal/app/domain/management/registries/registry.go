package registries

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
)

type Registry struct {
	gorm.Model
	BookID         uint `json:"bookId" gorm:"not null"`
	RegistryNumber int  `json:"registryNumber" gorm:"not null"`
}

func (brn Registry) TableName() string {
	return "book_registry_numbers"
}

type Filter struct {
	BookID         string `json:"bookId,omitempty" column:"book_registry_numbers.book_id" array:"false"`
	RegistryNumber string `json:"registryNumber,omitempty" column:"book_registry_numbers.registry_number" array:"false"`
}

func (f Filter) GetFieldNames() []string {
	return []string{"BookID", "RegistryNumber"}
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
