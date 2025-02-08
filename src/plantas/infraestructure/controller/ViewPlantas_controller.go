package controller

import (
	"net/http"

	"api/src/plantas/application"

	"github.com/gin-gonic/gin"
)

type PlantaViewController struct {
	useCase           *application.PlantaViewUseCase
	newRecordsUseCase *application.PlantaNewRecordsUseCase
}

func NewPlantaViewController(useCase *application.PlantaViewUseCase, newRecordsUseCase *application.PlantaNewRecordsUseCase) *PlantaViewController {
	return &PlantaViewController{
		useCase:           useCase,
		newRecordsUseCase: newRecordsUseCase,
	}
}

func (h *PlantaViewController) GetAll(c *gin.Context) {
	plantas, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plantas)
}

func (h *PlantaViewController) GetNewPlants(c *gin.Context) {
	lastUpdated := c.Query("last_updated")
	if lastUpdated == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere el parámetro last_updated"})
		return
	}

	plantas, err := h.newRecordsUseCase.GetNewRecords(lastUpdated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plantas)
}
