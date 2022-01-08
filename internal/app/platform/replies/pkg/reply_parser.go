package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ParserToReplyObj(obj interface{}) (replies.Reply, *errors.ApiError) {
	data, ok := obj.(*replies.Reply)
	if !ok {
		return replies.Reply{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}

func ParserToSliceReplyObj(obj interface{}) ([]replies.Reply, *errors.ApiError) {
	if obj == nil {
		return []replies.Reply{}, nil
	}
	data, ok := obj.(*[]replies.Reply)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParsetObj,
		}
	}
	return *data, nil
}
