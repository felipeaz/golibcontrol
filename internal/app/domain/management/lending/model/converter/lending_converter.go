package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
)

func ConvertToLendingObj(obj interface{}) (model.Lending, *errors.ApiError) {
	lendingObj, ok := obj.(*model.Lending)
	if !ok {
		return model.Lending{}, &errors.ApiError{
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
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *lendingObj, nil
}
