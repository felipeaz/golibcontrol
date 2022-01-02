package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	"github.com/gin-gonic/gin"
)

type ConferenceHandler struct {
	Module conferences.Module
}

func NewConferenceHandler(module conferences.Module) ConferenceHandler {
	return ConferenceHandler{
		Module: module,
	}
}

func (h ConferenceHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ConferenceHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ConferenceHandler) Create(c *gin.Context) {
	conference, apiError := pkg.AssociateConferenceInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(conference)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h ConferenceHandler) Update(c *gin.Context) {
	upConference, apiError := pkg.AssociateConferenceInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upConference)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ConferenceHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
