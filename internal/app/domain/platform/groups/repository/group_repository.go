package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/pkg"
)

type GroupRepository struct {
	DB database.GORMServiceInterface
}

func NewGroupRepository(db database.GORMServiceInterface) GroupRepository {
	return GroupRepository{
		DB: db,
	}
}

func (r GroupRepository) Get() ([]groups.Group, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]groups.Group{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceGroupObj(result)
}

func (r GroupRepository) Find(id string) (groups.Group, *errors.ApiError) {
	result, err := r.DB.Fetch(&groups.Group{}, id)
	if err != nil {
		return groups.Group{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToGroupObj(result)
}

func (r GroupRepository) Create(group groups.Group) (*groups.Group, *errors.ApiError) {
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

func (r GroupRepository) Update(id string, upGroup groups.Group) *errors.ApiError {
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
	err := r.DB.Remove(&groups.Group{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
