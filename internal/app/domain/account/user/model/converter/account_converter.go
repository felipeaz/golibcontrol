package converter

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/user/model"
)

func ConvertToAccountObj(obj interface{}) (model.Account, *errors.ApiError) {
	accountObj, ok := obj.(*model.Account)
	if !ok {
		return model.Account{}, &errors.ApiError{
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
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *accountObj, nil
}
