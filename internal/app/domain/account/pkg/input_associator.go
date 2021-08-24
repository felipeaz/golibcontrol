package pkg

import (
	"net/http"
	"os"

	accountModel "github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

// AssociateAccountInput is responsible of associating the params to the user model.
func AssociateAccountInput(c *gin.Context) (account accountModel.Account, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&account)
	if err != nil {
		return accountModel.Account{}, &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateSessionInput is responsible of associating the params to the user model.
func AssociateSessionInput(c *gin.Context) (session accountModel.UserSession, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&session)
	if err != nil {
		return accountModel.UserSession{}, &errors.ApiError{
			Service: os.Getenv("ACCOUNT_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
