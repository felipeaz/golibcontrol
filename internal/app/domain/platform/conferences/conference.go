package conferences

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

//const (
//	PRE_STATUS    = "PRE"
//	ONLINE_STATUS = "ON"
//	POST_STATUS   = "POST"
//)

type Conference struct {
	ID                    uint                    `json:"id" gorm:"primaryKey"`
	Subject               string                  `json:"subject" binding:"required"`
	StartDate             time.Time               `json:"startDate" binding:"required" time_format:"2006-01-02 15:04:05"`
	EndDate               time.Time               `json:"endDate" binding:"required" time_format:"2006-01-02 15:04:05"`
	Duration              int                     `json:"duration" binding:"required"`
	Status                string                  `json:"status" binding:"required"`
	MeetingHash           string                  `json:"meetingHash"`
	ConferenceSubscribers []ConferenceSubscribers `json:"subscribers"`
	CreatedAt             time.Time               `json:"createdAt"`
	UpdatedAt             time.Time               `json:"updatedAt"`
}

func (c Conference) TableName() string {
	return "conferences"
}

type Module interface {
	Get() ([]Conference, *errors.ApiError)
	Find(id string) (Conference, *errors.ApiError)
	Create(conference Conference) (*Conference, *errors.ApiError)
	Update(id string, upConference Conference) *errors.ApiError
	Delete(id string) *errors.ApiError
	Subscribe(subscription ConferenceSubscribers) (*ConferenceSubscribers, *errors.ApiError)
	Unsubscribe(subscription ConferenceSubscribers) *errors.ApiError
}

type Repository interface {
	Get() ([]Conference, *errors.ApiError)
	Find(id string) (Conference, *errors.ApiError)
	Create(conference Conference) (*Conference, *errors.ApiError)
	Update(id string, upConference Conference) *errors.ApiError
	Delete(id string) *errors.ApiError
	Subscribe(subscription ConferenceSubscribers) (*ConferenceSubscribers, *errors.ApiError)
	Unsubscribe(subscription ConferenceSubscribers) *errors.ApiError
}

type ConferenceSubscribers struct {
	ConferenceID uint `json:"conferenceId"`
	StudentID    uint `json:"studentId" binding:"required"`
}

func (s ConferenceSubscribers) TableName() string {
	return "conference_subscribers"
}
