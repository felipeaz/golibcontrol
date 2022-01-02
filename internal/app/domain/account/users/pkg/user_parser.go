package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ParseAccountEntry is responsible for associating the params to the user model.
func ParseAccountEntry(c *gin.Context) (account users.Account, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&account)
	if err != nil {
		return users.Account{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

func ParseInterfaceToAccount(obj interface{}) (users.Account, *errors.ApiError) {
	accountObj, ok := obj.(*users.Account)
	if !ok {
		return users.Account{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *accountObj, nil
}

func ParseInterfaceToSliceAccount(obj interface{}) ([]users.Account, *errors.ApiError) {
	if obj == nil {
		return []users.Account{}, nil
	}
	accountObj, ok := obj.(*[]users.Account)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *accountObj, nil
}
