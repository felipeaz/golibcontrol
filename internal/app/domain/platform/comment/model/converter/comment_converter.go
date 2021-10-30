package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model"
)

func ConvertToCommentObj(obj interface{}) (model.Comment, *errors.ApiError) {
	comment, ok := obj.(*model.Comment)
	if !ok {
		return model.Comment{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *comment, nil
}

func ConvertToSliceCommentObj(obj interface{}) ([]model.Comment, *errors.ApiError) {
	comments, ok := obj.(*[]model.Comment)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *comments, nil
}
