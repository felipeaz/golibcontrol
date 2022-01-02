package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	Module reviews.Module
}

func NewReviewHandler(module reviews.Module) ReviewHandler {
	return ReviewHandler{
		Module: module,
	}
}

func (h ReviewHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ReviewHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ReviewHandler) Create(c *gin.Context) {
	review, apiError := pkg.AssociateReviewInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(review)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h ReviewHandler) Update(c *gin.Context) {
	upReview, apiError := pkg.AssociateReviewInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upReview)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ReviewHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
