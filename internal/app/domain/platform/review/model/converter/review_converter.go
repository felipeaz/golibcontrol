package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model"
)

func ConvertToReviewObj(obj interface{}) (model.Review, *errors.ApiError) {
	review, ok := obj.(*model.Review)
	if !ok {
		return model.Review{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *review, nil
}

func ConvertToSliceReviewObj(obj interface{}) ([]model.Review, *errors.ApiError) {
	reviews, ok := obj.(*[]model.Review)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *reviews, nil
}
