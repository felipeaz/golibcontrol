package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/model"
)

func ConvertToReplyObj(obj interface{}) (model.Reply, *errors.ApiError) {
	comment, ok := obj.(*model.Reply)
	if !ok {
		return model.Reply{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *comment, nil
}

func ConvertToSliceReplyObj(obj interface{}) ([]model.Reply, *errors.ApiError) {
	if obj == nil {
		return []model.Reply{}, nil
	}
	comments, ok := obj.(*[]model.Reply)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *comments, nil
}
