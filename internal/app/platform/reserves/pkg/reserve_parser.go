package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToReserveObj(obj interface{}) (reserves.Reserve, *errors.ApiError) {
	data, ok := obj.(*reserves.Reserve)
	if !ok {
		return reserves.Reserve{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
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
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}
