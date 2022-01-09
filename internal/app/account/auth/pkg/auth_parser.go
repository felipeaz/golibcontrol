package pkg

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ParseSessionEntry is responsible for associating the params to the user model.
func ParseSessionEntry(c *gin.Context) (session auth.Session, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&session)
	if err != nil {
		return auth.Session{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}
	return
}
