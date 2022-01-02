package groups

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"time"
)

type Group struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	MeetingHash string    `json:"meetingHash"`
	CreatedAt   time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `time_format:"2006-01-02 15:04:05"`
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
