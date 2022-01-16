package groups

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

type Group struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	MeetingHash string    `json:"meetingHash"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (g Group) TableName() string {
	return "groups"
}

type Module interface {
	Get() ([]Group, *errors.ApiError)
	Find(id string) (Group, *errors.ApiError)
	Create(group Group) (*Group, *errors.ApiError)
	Update(id string, upGroup Group) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type GroupRepositoryInterface interface {
	Get() ([]Group, *errors.ApiError)
	Find(id string) (Group, *errors.ApiError)
	Create(group Group) (*Group, *errors.ApiError)
	Update(id string, upGroup Group) *errors.ApiError
	Delete(id string) *errors.ApiError
}
