package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model"
)

func ConvertToReviewObj(obj interface{}) (model.Review, *errors.ApiError) {
	review, ok := obj.(*model.Review)
	if !ok {
		return model.Review{}, &errors.ApiError{
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
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
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *reviews, nil
}
