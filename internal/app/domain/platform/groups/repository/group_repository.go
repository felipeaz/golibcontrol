package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/model/converter"
)

type GroupRepository struct {
	DB database.GORMServiceInterface
}

func NewGroupRepository(db database.GORMServiceInterface) GroupRepository {
	return GroupRepository{
		DB: db,
	}
}

func (r GroupRepository) Get() ([]model.Group, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]model.Group{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	groups, apiError := converter.ConvertToSliceGroupObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return groups, nil
}

func (r GroupRepository) Find(id string) (model.Group, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Group{}, id)
	if err != nil {
		return model.Group{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	group, apiError := converter.ConvertToGroupObj(result)
	if apiError != nil {
		return model.Group{}, apiError
	}
	return group, nil
}

func (r GroupRepository) Create(group model.Group) (*model.Group, *errors.ApiError) {
	err := r.DB.Persist(&group)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &group, nil
}

func (r GroupRepository) Update(id string, upGroup model.Group) *errors.ApiError {
	err := r.DB.Refresh(&upGroup, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r GroupRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Group{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
