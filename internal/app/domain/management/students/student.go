package students

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

// Student contains all Student's table properties.
type Student struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Grade     string    `json:"grade"`
	Birthday  string    `json:"birthday"`
	AccountId uint      `json:"accountId" gorm:"index;unique;not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s Student) TableName() string {
	return "students"
}

func (s Student) GetFirstName() string {
	nameArr := strings.Split(s.Name, " ")
	return nameArr[0]
}

func (s Student) GetLastName() string {
	nameArr := strings.Split(s.Name, " ")
	return nameArr[len(nameArr)-1]
}

type Account struct {
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Phone          string `json:"phone"`
	StudentAccount bool   `json:"studentAccount" gorm:"<-:create;default:false"`
}

func (a *Account) BeforeUpdate(tx *gorm.DB) error {
	resp := tx.Model(&a).Update("student_account", a.StudentAccount)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

type AccountResponse struct {
	ID uint `json:"id"`
}

type Module interface {
	Get() ([]Student, *errors.ApiError)
	Find(id string) (Student, *errors.ApiError)
	Create(student Student, accountHost, accountRoute, tokenName, tokenValue string) (*Student, *errors.ApiError)
	Update(id string, upStudent Student) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() (students []Student, apiError *errors.ApiError)
	Find(id string) (student Student, apiError *errors.ApiError)
	Create(student Student) (*Student, *errors.ApiError)
	Update(id string, upStudent Student) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
