package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToReviewObj(obj interface{}) (reviews.Review, *errors.ApiError) {
	data, ok := obj.(*reviews.Review)
	if !ok {
		return reviews.Review{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}

func ParseToSliceReviewObj(obj interface{}) ([]reviews.Review, *errors.ApiError) {
	if obj == nil {
		return []reviews.Review{}, nil
	}
	data, ok := obj.(*[]reviews.Review)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}
