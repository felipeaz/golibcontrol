package pkg

import (
	commentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments"
	conferenceModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences"
	groupModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
	replyModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	reserveModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	reviewModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
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

// AssociateConferenceInput is responsible for associating the params to the user model.
func AssociateConferenceInput(c *gin.Context) (conference conferenceModel.Conference, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&conference)
	if err != nil {
		return conferenceModel.Conference{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateGroupInput is responsible for associating the params to the user model.
func AssociateGroupInput(c *gin.Context) (group groupModel.Group, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&group)
	if err != nil {
		return groupModel.Group{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
