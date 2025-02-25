package controller

import (
	"net/http"

	"api/src/plantas/application"

	"github.com/gin-gonic/gin"
)

type PlantaViewController struct {
	useCase *application.PlantaViewUseCase
}

func NewPlantaViewController(useCase *application.PlantaViewUseCase) *PlantaViewController {
	return &PlantaViewController{
		useCase: useCase,
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
