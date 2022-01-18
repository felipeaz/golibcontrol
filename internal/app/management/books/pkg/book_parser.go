package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"net/http"
)

func ParseToBookObj(obj interface{}) (books.Book, *errors.ApiError) {
	data, ok := obj.(*books.Book)
	if !ok {
		return books.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}

func ParseToSliceBookObj(obj interface{}) ([]books.Book, *errors.ApiError) {
	if obj == nil {
		return []books.Book{}, nil
	}
	data, ok := obj.(*[]books.Book)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToParseObj,
		}
	}
	return *data, nil
}
