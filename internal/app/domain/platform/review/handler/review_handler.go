package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/module/interface"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	Module _interface.ReviewModuleInterface
}

func NewReviewHandler(module _interface.ReviewModuleInterface) ReviewHandler {
	return ReviewHandler{
		Module: module,
	}
}

func (h ReviewHandler) Get(c *gin.Context) {
	reviews, apiError := h.Module.Get(c.Param("bookId"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func (h ReviewHandler) Find(c *gin.Context) {
	review, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

func (h ReviewHandler) Create(c *gin.Context) {
	review, apiError := pkg.AssociateReviewInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	id, apiError := h.Module.Create(review)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
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
