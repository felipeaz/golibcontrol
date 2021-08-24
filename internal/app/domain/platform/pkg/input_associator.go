package pkg

import (
	"net/http"
	"os"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	commentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model"
	reserveModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/model"
	reviewModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model"
	"github.com/gin-gonic/gin"
)

// AssociateCommentInput is responsible of associating the params to the user model.
func AssociateCommentInput(c *gin.Context) (comment commentModel.Comment, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		return commentModel.Comment{}, &errors.ApiError{
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateReserveInput is responsible of associating the params to the user model.
func AssociateReserveInput(c *gin.Context) (reserve reserveModel.Reserve, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&reserve)
	if err != nil {
		return reserveModel.Reserve{}, &errors.ApiError{
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateReviewInput is responsible of associating the params to the user model.
func AssociateReviewInput(c *gin.Context) (review reviewModel.Review, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&review)
	if err != nil {
		return reviewModel.Review{}, &errors.ApiError{
			Service: os.Getenv("PLATFORM_SERVICE_NAME"),
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
