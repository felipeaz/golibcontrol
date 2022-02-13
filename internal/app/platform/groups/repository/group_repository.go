package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/groups/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
	tx := r.DB.Preload("GroupSubscribers")
	result, err := r.DB.Find(tx, &[]groups.Group{})
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
	tx := r.DB.Preload("GroupSubscribers")
	tx = r.DB.Where(tx, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &groups.Group{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upGroup)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &groups.Group{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r GroupRepository) Subscribe(subscription groups.GroupSubscribers) (*groups.GroupSubscribers, *errors.ApiError) {
	err := r.DB.Persist(&subscription)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &subscription, nil
}

func (r GroupRepository) Unsubscribe(subscription groups.GroupSubscribers) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("student_id = %d AND group_id = %d",
		subscription.StudentID, subscription.GroupID))
	err := r.DB.Remove(tx, &groups.GroupSubscribers{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
