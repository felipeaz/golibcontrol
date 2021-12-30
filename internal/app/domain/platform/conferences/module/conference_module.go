package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type ConferenceModule struct {
	Repository _interface.ConferenceRepositoryInterface
	Log        logger.LogInterface
}

func NewConferenceModule(repo _interface.ConferenceRepositoryInterface, log logger.LogInterface) ConferenceModule {
	return ConferenceModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ConferenceModule) Get() ([]model.Conference, *errors.ApiError) {
	return m.Repository.Get()
}

func (m ConferenceModule) Find(id string) (model.Conference, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ConferenceModule) Create(comment model.Conference) (*model.Conference, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ConferenceModule) Update(id string, upConference model.Conference) *errors.ApiError {
	return m.Repository.Update(id, upConference)
}

func (m ConferenceModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
