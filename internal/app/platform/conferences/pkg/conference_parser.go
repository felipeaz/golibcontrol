package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToConferenceObj(obj interface{}) (conferences.Conference, *errors.ApiError) {
	data, ok := obj.(*conferences.Conference)
	if !ok {
		return conferences.Conference{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}

func ParseToSliceConferenceObj(obj interface{}) ([]conferences.Conference, *errors.ApiError) {
	if obj == nil {
		return []conferences.Conference{}, nil
	}
	data, ok := obj.(*[]conferences.Conference)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}
