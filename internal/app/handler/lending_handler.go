package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
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

}

// Find return one lending by ID.
func (h LendingHandler) Find(c *gin.Context) {

}

// Create creates a lending.
func (h LendingHandler) Create(c *gin.Context) {

}

// Update update an existent lending.
func (h LendingHandler) Update(c *gin.Context) {

}

// Delete delete an existent lending by id.
func (h LendingHandler) Delete(c *gin.Context) {

}
