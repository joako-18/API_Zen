package controllers

import (
	"api/src/viveros/application"
	"api/src/viveros/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateViveroController struct {
	useCase *application.CreateViveroUseCase
}

func NewCreateViveroController(useCase *application.CreateViveroUseCase) *CreateViveroController {
	return &CreateViveroController{useCase: useCase}
}

func (h *CreateViveroController) Create(c *gin.Context) {
	var vivero entities.Vivero
	if err := c.ShouldBindJSON(&vivero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.useCase.Create(vivero)
	c.JSON(http.StatusCreated, vivero)
}
