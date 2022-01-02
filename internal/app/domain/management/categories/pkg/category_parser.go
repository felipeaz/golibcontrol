package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ParseToCategoryObj(obj interface{}) (categories.Category, *errors.ApiError) {
	categoryObj, ok := obj.(*categories.Category)
	if !ok {
		return categories.Category{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *categoryObj, nil
}

func ParseToSliceCategoryObj(obj interface{}) ([]categories.Category, *errors.ApiError) {
	if obj == nil {
		return []categories.Category{}, nil
	}
	categoryObj, ok := obj.(*[]categories.Category)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *categoryObj, nil
}
