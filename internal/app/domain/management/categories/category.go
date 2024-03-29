package categories

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

// Category contains all Category's table properties.
type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
