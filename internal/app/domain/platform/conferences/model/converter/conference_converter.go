package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences/model"
)

func ConvertToConferenceObj(obj interface{}) (model.Conference, *errors.ApiError) {
	conference, ok := obj.(*model.Conference)
	if !ok {
		return model.Conference{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *conference, nil
}

func ConvertToSliceConferenceObj(obj interface{}) ([]model.Conference, *errors.ApiError) {
	if obj == nil {
		return []model.Conference{}, nil
	}
	conferences, ok := obj.(*[]model.Conference)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *conferences, nil
}
