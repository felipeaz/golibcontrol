package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ParseToCategoryObj(obj interface{}) (categories.Category, *errors.ApiError) {
	data, ok := obj.(*categories.Category)
	if !ok {
		return categories.Category{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}

func ParseToSliceCategoryObj(obj interface{}) ([]categories.Category, *errors.ApiError) {
	if obj == nil {
		return []categories.Category{}, nil
	}
	data, ok := obj.(*[]categories.Category)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}
