package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ParseToReserveObj(obj interface{}) (reserves.Reserve, *errors.ApiError) {
	data, ok := obj.(*reserves.Reserve)
	if !ok {
		return reserves.Reserve{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}

func ParseToSliceReserveObj(obj interface{}) ([]reserves.Reserve, *errors.ApiError) {
	if obj == nil {
		return []reserves.Reserve{}, nil
	}
	data, ok := obj.(*[]reserves.Reserve)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}
