package controller

import (
	"net/http"

	"api/src/plantas/application"
	"api/src/plantas/domain/entities"

	"github.com/gin-gonic/gin"
)

type PlantaUpdateController struct {
	useCase *application.PlantaUpdateUseCase
}

func NewPlantaUpdateController(useCase *application.PlantaUpdateUseCase) *PlantaUpdateController {
	return &PlantaUpdateController{useCase: useCase}
}

func (h *PlantaUpdateController) Update(c *gin.Context) {
	var planta entities.Planta
	if err := c.ShouldBindJSON(&planta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.Update(planta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, planta)
}
