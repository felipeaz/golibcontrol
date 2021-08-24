package converter

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
)

func ConvertToAccountObj(obj interface{}) (model.Account, *errors.ApiError) {
	accountObj, ok := obj.(*model.Account)
	if !ok {
		return model.Account{}, &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *accountObj, nil
}

func ConvertToSliceAccountObj(obj interface{}) ([]model.Account, *errors.ApiError) {
	accountObj, ok := obj.(*[]model.Account)
	if !ok {
		return nil, &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *accountObj, nil
}
