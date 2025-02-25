package controller

import (
	"net/http"
	"replica/src/planta/application"

	"github.com/gin-gonic/gin"
)

type PlantaShortPollingController struct {
	useCase *application.PlantaShortPollingUseCase
}

func NewPlantaShortPollingController(useCase *application.PlantaShortPollingUseCase) *PlantaShortPollingController {
	return &PlantaShortPollingController{useCase: useCase}
}

func (h *PlantaShortPollingController) GetAll(c *gin.Context) {
	plantas, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo las plantas"})
		return
	}
	c.JSON(http.StatusOK, plantas)
}
