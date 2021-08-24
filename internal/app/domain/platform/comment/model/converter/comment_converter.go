package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model"
)

func ConvertToCommentObj(obj interface{}) (model.Comment, *errors.ApiError) {
	comment, ok := obj.(*model.Comment)
	if !ok {
		return model.Comment{}, &errors.ApiError{
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
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
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *comments, nil
}
