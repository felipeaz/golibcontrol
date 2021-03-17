package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
)

// CategoryHandler handle the category router call.
type CategoryHandler struct {
	Module module.CategoryModule
}

// NewCategoryHandler returns an instance of category handler.
func NewCategoryHandler(db *gorm.DB) CategoryHandler {
	return CategoryHandler{
		Module: module.CategoryModule{
			Repository: repository.CategoryRepository{
				DB: db,
			},
		},
	}
}

// Get returns all categories.
func (h CategoryHandler) Get(c *gin.Context) {
	categories, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// Find return one category by ID.
func (h CategoryHandler) Find(c *gin.Context) {
	category, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Create persist a category to the database.
func (h CategoryHandler) Create(c *gin.Context) {
	category, err := pkg.AssociateCategoryInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	id, apiError := h.Module.Create(category)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Update update an existent category.
func (h CategoryHandler) Update(c *gin.Context) {
	upCategory, err := pkg.AssociateCategoryInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	category, apiError := h.Module.Update(c.Param("id"), upCategory)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Delete delete an existent category.
func (h CategoryHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
