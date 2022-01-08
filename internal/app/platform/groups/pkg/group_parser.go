package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ParseToGroupObj(obj interface{}) (groups.Group, *errors.ApiError) {
	data, ok := obj.(*groups.Group)
	if !ok {
		return groups.Group{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}

func ParseToSliceGroupObj(obj interface{}) ([]groups.Group, *errors.ApiError) {
	if obj == nil {
		return []groups.Group{}, nil
	}
	data, ok := obj.(*[]groups.Group)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}
