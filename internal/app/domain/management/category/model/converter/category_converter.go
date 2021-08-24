package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/model"
)

func ConvertToCategoryObj(obj interface{}) (model.Category, *errors.ApiError) {
	categoryObj, ok := obj.(*model.Category)
	if !ok {
		return model.Category{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *categoryObj, nil
}

func ConvertToSliceCategoryObj(obj interface{}) ([]model.Category, *errors.ApiError) {
	categoryObj, ok := obj.(*[]model.Category)
	if !ok {
		return nil, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *categoryObj, nil
}
