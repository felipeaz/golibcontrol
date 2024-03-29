package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

type ConferenceModule struct {
	Repository conferences.Repository
	Log        logger.LogInterface
}

func NewConferenceModule(repo conferences.Repository, log logger.LogInterface) ConferenceModule {
	return ConferenceModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ConferenceModule) Get() ([]conferences.Conference, *errors.ApiError) {
	return m.Repository.Get()
}

func (m ConferenceModule) Find(id string) (conferences.Conference, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ConferenceModule) Create(conference conferences.Conference) (*conferences.Conference, *errors.ApiError) {
	return m.Repository.Create(conference)
}

func (m ConferenceModule) Update(id string, upConference conferences.Conference) *errors.ApiError {
	return m.Repository.Update(id, upConference)
}

func (m ConferenceModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}

func (m ConferenceModule) Subscribe(subscription conferences.ConferenceSubscribers) (*conferences.ConferenceSubscribers, *errors.ApiError) {
	return m.Repository.Subscribe(subscription)
}

func (m ConferenceModule) Unsubscribe(subscription conferences.ConferenceSubscribers) *errors.ApiError {
	return m.Repository.Unsubscribe(subscription)
}
