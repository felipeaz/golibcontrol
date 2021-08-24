package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
)

func ConvertToLendingObj(obj interface{}) (model.Lending, *errors.ApiError) {
	lendingObj, ok := obj.(*model.Lending)
	if !ok {
		return model.Lending{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *lendingObj, nil
}

func ConvertToSliceLendingObj(obj interface{}) ([]model.Lending, *errors.ApiError) {
	lendingObj, ok := obj.(*[]model.Lending)
	if !ok {
		return nil, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *lendingObj, nil
}
