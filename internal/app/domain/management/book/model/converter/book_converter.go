package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/model"
)

func ConvertToBookObj(obj interface{}) (model.Book, *errors.ApiError) {
	bookObj, ok := obj.(*model.Book)
	if !ok {
		return model.Book{}, &errors.ApiError{
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
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
			Service: os.Getenv("MANAGEMENT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *bookObj, nil
}
