package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/book/model"
)

func ConvertToBookObj(obj interface{}) (model.Book, *errors.ApiError) {
	bookObj, ok := obj.(*model.Book)
	if !ok {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *bookObj, nil
}

func ConvertToSliceBookObj(obj interface{}) ([]model.Book, *errors.ApiError) {
	bookObj, ok := obj.(*[]model.Book)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *bookObj, nil
}
