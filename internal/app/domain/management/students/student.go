package students

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"strings"
	"time"
)

// Student contains all Student's table properties.
type Student struct {
	ID        string    `json:"id" binding:"required" gorm:"primaryKey;unique;not null"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Grade     string    `json:"grade" binding:"required"`
	Birthday  string    `json:"birthday" binding:"required"`
	AccountId uint      `json:"accountId" gorm:"index;unique;not null"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
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
	StudentAccount bool   `json:"studentAccount" gorm:"<-:create"`
}

type AccountResponse struct {
	ID uint `json:"id"`
}
