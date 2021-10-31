package pkg

import (
	"net/http"

	accountModel "github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

// AssociateAccountInput is responsible for associating the params to the user model.
func AssociateAccountInput(c *gin.Context) (account accountModel.Account, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&account)
	if err != nil {
		return accountModel.Account{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateSessionInput is responsible for associating the params to the user model.
func AssociateSessionInput(c *gin.Context) (session accountModel.UserSession, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&session)
	if err != nil {
		return accountModel.UserSession{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
