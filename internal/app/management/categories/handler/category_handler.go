package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/app/management/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryHandler handle the category router call.
type CategoryHandler struct {
	Module categories.Module
}

// NewCategoryHandler returns an instance of category handler.
func NewCategoryHandler(module categories.Module) CategoryHandler {
	return CategoryHandler{
		Module: module,
	}
}

// Get returns all categories.
func (h CategoryHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Find return one category by ID.
func (h CategoryHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create persist a category to the database.
func (h CategoryHandler) Create(c *gin.Context) {
	category, apiError := pkg.AssociateCategoryInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(category)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update update an existent category.
func (h CategoryHandler) Update(c *gin.Context) {
	upCategory, apiError := pkg.AssociateCategoryInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upCategory)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}

// Delete delete an existent category.
func (h CategoryHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}
