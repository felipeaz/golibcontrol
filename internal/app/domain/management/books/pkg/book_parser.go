package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func ParseToBookObj(obj interface{}) (books.Book, *errors.ApiError) {
	bookObj, ok := obj.(*books.Book)
	if !ok {
		return books.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *bookObj, nil
}

func ParseToSliceBookObj(obj interface{}) ([]books.Book, *errors.ApiError) {
	if obj == nil {
		return []books.Book{}, nil
	}
	bookObj, ok := obj.(*[]books.Book)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *bookObj, nil
}
