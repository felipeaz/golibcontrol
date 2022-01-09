package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
	result, err := r.DB.Find(nil, &[]conferences.Conference{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &conferences.Conference{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upConference)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &conferences.Conference{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
