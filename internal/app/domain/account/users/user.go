package users

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
)

// Account contains all Account's table properties.
type Account struct {
	gorm.Model
	ConsumerId     string `json:"consumerId" gorm:"not null"`
	Email          string `json:"email" binding:"required" gorm:"unique"`
	Password       string `json:"password" binding:"required"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Phone          string `json:"phone"`
	StudentAccount bool   `json:"studentAccount" gorm:"<-:create"`
}

func (a Account) TableName() string {
	return "accounts"
}

type Module interface {
	Get() ([]Account, *errors.ApiError)
	Find(id string) (Account, *errors.ApiError)
	Create(account Account) (*Account, *errors.ApiError)
	Update(id string, upAccount Account) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() (accounts []Account, apiError *errors.ApiError)
	Find(id string) (account Account, apiError *errors.ApiError)
	FindWhere(fieldName, fieldValue string) (account Account, apiError *errors.ApiError)
	Create(account Account) (*Account, *errors.ApiError)
	Update(id string, upAccount Account) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
