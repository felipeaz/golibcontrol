package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/model"
)

func ConvertToGroupObj(obj interface{}) (model.Group, *errors.ApiError) {
	group, ok := obj.(*model.Group)
	if !ok {
		return model.Group{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *group, nil
}

func ConvertToSliceGroupObj(obj interface{}) ([]model.Group, *errors.ApiError) {
	if obj == nil {
		return []model.Group{}, nil
	}
	groups, ok := obj.(*[]model.Group)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *groups, nil
}
