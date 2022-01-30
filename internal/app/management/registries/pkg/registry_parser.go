package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToRegistryObj(obj interface{}) (registries.Registry, *errors.ApiError) {
	data, ok := obj.(*registries.Registry)
	if !ok {
		return registries.Registry{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}

func ParseToSliceRegistryObj(obj interface{}) ([]registries.Registry, *errors.ApiError) {
	if obj == nil {
		return []registries.Registry{}, nil
	}
	data, ok := obj.(*[]registries.Registry)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}
