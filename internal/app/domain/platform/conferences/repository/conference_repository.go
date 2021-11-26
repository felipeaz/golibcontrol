package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences/model/converter"
)

type ConferenceRepository struct {
	DB database.GORMServiceInterface
}

func NewConferenceRepository(db database.GORMServiceInterface) ConferenceRepository {
	return ConferenceRepository{
		DB: db,
	}
}

func (r ConferenceRepository) Get() ([]model.Conference, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]model.Conference{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	conferences, apiError := converter.ConvertToSliceConferenceObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return conferences, nil
}

func (r ConferenceRepository) Find(id string) (model.Conference, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Conference{}, id)
	if err != nil {
		return model.Conference{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	conference, apiError := converter.ConvertToConferenceObj(result)
	if apiError != nil {
		return model.Conference{}, apiError
	}
	return conference, nil
}

func (r ConferenceRepository) Create(conference model.Conference) (uint, *errors.ApiError) {
	err := r.DB.Persist(&conference)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return conference.ID, nil
}

func (r ConferenceRepository) Update(id string, upConference model.Conference) *errors.ApiError {
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
	err := r.DB.Remove(&model.Conference{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
