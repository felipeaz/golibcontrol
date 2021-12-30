package pkg

import (
	session_model "github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth/model"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/model"
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

// AssociateAccountInput is responsible for associating the params to the user model.
func AssociateAccountInput(c *gin.Context) (account model.Account, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&account)
	if err != nil {
		return model.Account{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateSessionInput is responsible for associating the params to the user model.
func AssociateSessionInput(c *gin.Context) (session session_model.UserSession, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&session)
	if err != nil {
		return session_model.UserSession{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
