package users

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
	"time"
)

// Account contains all Account's table properties.
type Account struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ConsumerId     string    `json:"consumerId" gorm:"not null"`
	Email          string    `json:"email" binding:"required" gorm:"unique"`
	Password       string    `json:"password" binding:"required"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Phone          string    `json:"phone"`
	StudentAccount bool      `json:"studentAccount" gorm:"<-:create;default:false"`
}

func (a Account) TableName() string {
	return "accounts"
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	var err error
	crypto := &pkg.Cryptor{}
	a.Password, err = crypto.EncryptPassword(a.Password)
	if err != nil {
		return err
	}
	return nil
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
