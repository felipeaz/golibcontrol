package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/model"
)

func ConvertToReserveObj(obj interface{}) (model.Reserve, *errors.ApiError) {
	comment, ok := obj.(*model.Reserve)
	if !ok {
		return model.Reserve{}, &errors.ApiError{
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *comment, nil
}

func ConvertToSliceReserveObj(obj interface{}) ([]model.Reserve, *errors.ApiError) {
	comments, ok := obj.(*[]model.Reserve)
	if !ok {
		return nil, &errors.ApiError{
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *comments, nil
}
