package groups

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

type Group struct {
	ID               uint               `json:"id" gorm:"primaryKey"`
	Name             string             `json:"name" binding:"required"`
	Description      string             `json:"description" binding:"required"`
	Status           string             `json:"status" binding:"required"`
	MeetingHash      string             `json:"meetingHash"`
	GroupSubscribers []GroupSubscribers `json:"subscribers"`
	CreatedAt        time.Time          `json:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt"`
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
	Subscribe(subscription GroupSubscribers) (*GroupSubscribers, *errors.ApiError)
	Unsubscribe(subscription GroupSubscribers) *errors.ApiError
}

type GroupRepositoryInterface interface {
	Get() ([]Group, *errors.ApiError)
	Find(id string) (Group, *errors.ApiError)
	Create(group Group) (*Group, *errors.ApiError)
	Update(id string, upGroup Group) *errors.ApiError
	Delete(id string) *errors.ApiError
	Subscribe(subscription GroupSubscribers) (*GroupSubscribers, *errors.ApiError)
	Unsubscribe(subscription GroupSubscribers) *errors.ApiError
}

type GroupSubscribers struct {
	GroupID   uint `json:"groupId"`
	StudentID uint `json:"studentId" binding:"required"`
}

func (s GroupSubscribers) TableName() string {
	return "group_subscribers"
}
