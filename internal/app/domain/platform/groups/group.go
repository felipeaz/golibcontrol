package groups

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

type Group struct {
	ID                 uint               `json:"id" gorm:"primaryKey"`
	Name               string             `json:"name" binding:"required"`
	Description        string             `json:"description" binding:"required"`
	Status             string             `json:"status" binding:"required"`
	MeetingHash        string             `json:"meetingHash"`
	MeetingDescription string             `json:"meetingDescription"`
	GroupSubscribers   []GroupSubscribers `json:"subscribers"`
	CreatedAt          time.Time          `json:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt"`
}

func (g Group) TableName() string {
	return "groups"
}

type Filter struct {
	StudentID string `form:"studentId,omitempty" column:"group_subscribers.student_id" array:"false" like:"false"`
}

func (f Filter) GetFieldNames() []string {
	return []string{"StudentID"}
}

type Module interface {
	Get() ([]Group, *errors.ApiError)
	GetByFilter(filter Filter) ([]Group, *errors.ApiError)
	Find(id string) (Group, *errors.ApiError)
	Create(group Group) (*Group, *errors.ApiError)
	Update(id string, upGroup Group) *errors.ApiError
	Delete(id string) *errors.ApiError
	Subscribe(subscription GroupSubscribers) (*GroupSubscribers, *errors.ApiError)
	Unsubscribe(subscription GroupSubscribers) *errors.ApiError
}

type GroupRepositoryInterface interface {
	Get() ([]Group, *errors.ApiError)
	GetByFilter(filter Filter) ([]Group, *errors.ApiError)
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
