package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/lending/module"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/lending/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/pkg"
)

// LendingHandler handle the lending router call.
type LendingHandler struct {
	Module module.LendingModule
}

// NewLendingHandler Return an instance of this handler.
func NewLendingHandler(db *gorm.DB) LendingHandler {
	return LendingHandler{
		Module: module.LendingModule{
			Repository: repository.LendingRepository{
				DB: db,
			},
		},
	}
}

// Get returns all lendings.
func (h LendingHandler) Get(c *gin.Context) {
	lendings, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lendings})
}

// Find return one lending by ID.
func (h LendingHandler) Find(c *gin.Context) {
	lending, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lending})
}

// Create persist a lending to the database.
func (h LendingHandler) Create(c *gin.Context) {
	lending, err := pkg.AssociateLendingInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	id, apiError := h.Module.Create(lending)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Update update an existent lending.
func (h LendingHandler) Update(c *gin.Context) {
	upLending, err := pkg.AssociateLendingInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	lending, apiError := h.Module.Update(c.Param("id"), upLending)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lending})
}

// Delete delete an existent lending.
func (h LendingHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
