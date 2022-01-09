package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToLendingObj(obj interface{}) (lending.Lending, *errors.ApiError) {
	data, ok := obj.(*lending.Lending)
	if !ok {
		return lending.Lending{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}

func ParseToSliceLendingObj(obj interface{}) ([]lending.Lending, *errors.ApiError) {
	if obj == nil {
		return []lending.Lending{}, nil
	}
	data, ok := obj.(*[]lending.Lending)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}
