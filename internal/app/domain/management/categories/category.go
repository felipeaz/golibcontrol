package categories

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"gorm.io/gorm"
)

// Category contains all Category's table properties.
type Category struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (c Category) TableName() string {
	return "categories"
}

type Module interface {
	Get() ([]Category, *errors.ApiError)
	Find(id string) (Category, *errors.ApiError)
	Create(category Category) (*Category, *errors.ApiError)
	Update(id string, upCategory Category) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() (categories []Category, apiError *errors.ApiError)
	Find(id string) (category Category, apiError *errors.ApiError)
	Create(category Category) (*Category, *errors.ApiError)
	Update(id string, upCategory Category) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
