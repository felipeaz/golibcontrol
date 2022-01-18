package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToCategoryObj(obj interface{}) (categories.Category, *errors.ApiError) {
	data, ok := obj.(*categories.Category)
	if !ok {
		return categories.Category{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
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
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}
