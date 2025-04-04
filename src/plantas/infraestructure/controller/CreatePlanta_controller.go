package controller

import (
	"api/src/plantas/application"
	"api/src/plantas/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlantaCreateController struct {
	useCase *application.PlantaCreateUseCase
}

func NewPlantaCreateController(useCase *application.PlantaCreateUseCase) *PlantaCreateController {
	return &PlantaCreateController{useCase: useCase}
}

func (h *PlantaCreateController) Create(c *gin.Context) {
	var planta entities.Planta
	if err := c.ShouldBindJSON(&planta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.Create(planta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, planta)
}
