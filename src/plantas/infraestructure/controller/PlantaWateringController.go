package controller

import (
	"net/http"
	"strconv"

	"api/src/plantas/application"

	"github.com/gin-gonic/gin"
)

type PlantaWateringController struct {
	useCase *application.PlantaWateringUseCase
}

func NewPlantaWateringController(useCase *application.PlantaWateringUseCase) *PlantaWateringController {
	return &PlantaWateringController{useCase: useCase}
}

func (h *PlantaWateringController) Water(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.useCase.Water(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mensaje de riego enviado"})
}
