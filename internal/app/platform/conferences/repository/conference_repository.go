package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/pkg"
)

type ConferenceRepository struct {
	DB database.GORMServiceInterface
}

func NewConferenceRepository(db database.GORMServiceInterface) ConferenceRepository {
	return ConferenceRepository{
		DB: db,
	}
}

func (r ConferenceRepository) Get() ([]conferences.Conference, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]conferences.Conference{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceConferenceObj(result)
}

func (r ConferenceRepository) Find(id string) (conferences.Conference, *errors.ApiError) {
	result, err := r.DB.Fetch(&conferences.Conference{}, id)
	if err != nil {
		return conferences.Conference{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToConferenceObj(result)
}

func (r ConferenceRepository) Create(conference conferences.Conference) (*conferences.Conference, *errors.ApiError) {
	err := r.DB.Persist(&conference)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &conference, nil
}

func (r ConferenceRepository) Update(id string, upConference conferences.Conference) *errors.ApiError {
	err := r.DB.Refresh(&upConference, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r ConferenceRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&conferences.Conference{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
