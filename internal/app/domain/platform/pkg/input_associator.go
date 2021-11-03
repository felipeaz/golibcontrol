package pkg

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	commentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/model"
	replyModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/model"
	reserveModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/model"
	reviewModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/model"
	"github.com/gin-gonic/gin"
)

// AssociateCommentInput is responsible for associating the params to the user model.
func AssociateCommentInput(c *gin.Context) (comment commentModel.Comment, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		return commentModel.Comment{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateReplyInput is responsible for associating the params to the user model.
func AssociateReplyInput(c *gin.Context) (reply replyModel.Reply, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&reply)
	if err != nil {
		return replyModel.Reply{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateReserveInput is responsible for associating the params to the user model.
func AssociateReserveInput(c *gin.Context) (reserve reserveModel.Reserve, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&reserve)
	if err != nil {
		return reserveModel.Reserve{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateReviewInput is responsible for associating the params to the user model.
func AssociateReviewInput(c *gin.Context) (review reviewModel.Review, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&review)
	if err != nil {
		return reviewModel.Review{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
