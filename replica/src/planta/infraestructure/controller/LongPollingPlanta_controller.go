package controller

import (
	"net/http"
	"replica/src/planta/application"

	"github.com/gin-gonic/gin"
)

type PlantaLongPollingController struct {
	useCase *application.PlantaLongPollingUseCase
}

func NewPlantaLongPollingController(useCase *application.PlantaLongPollingUseCase) *PlantaLongPollingController {
	return &PlantaLongPollingController{useCase: useCase}
}

func (h *PlantaLongPollingController) LongPolling(c *gin.Context) {
	plantas, err := h.useCase.WaitForChanges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error esperando cambios"})
		return
	}
	c.JSON(http.StatusOK, plantas)
}
