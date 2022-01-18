package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParserToReplyObj(obj interface{}) (replies.Reply, *errors.ApiError) {
	data, ok := obj.(*replies.Reply)
	if !ok {
		return replies.Reply{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
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
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}
