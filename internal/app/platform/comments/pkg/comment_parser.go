package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToCommentObj(obj interface{}) (comments.Comment, *errors.ApiError) {
	data, ok := obj.(*comments.Comment)
	if !ok {
		return comments.Comment{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}

func ParseToSliceCommentObj(obj interface{}) ([]comments.Comment, *errors.ApiError) {
	if obj == nil {
		return []comments.Comment{}, nil
	}
	data, ok := obj.(*[]comments.Comment)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}
