package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences/model"
)

type ConferenceRepositoryInterface interface {
	Get() ([]model.Conference, *errors.ApiError)
	Find(id string) (model.Conference, *errors.ApiError)
	Create(conference model.Conference) (uint, *errors.ApiError)
	Update(id string, upConference model.Conference) *errors.ApiError
	Delete(id string) *errors.ApiError
}
