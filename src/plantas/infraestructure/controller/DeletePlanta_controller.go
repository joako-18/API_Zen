package controller

import (
	"net/http"
	"strconv"

	"api/src/plantas/application"

	"github.com/gin-gonic/gin"
)

type PlantaDeleteController struct {
	useCase *application.PlantaDeleteUseCase
}

func NewPlantaDeleteController(useCase *application.PlantaDeleteUseCase) *PlantaDeleteController {
	return &PlantaDeleteController{useCase: useCase}
}

func (h *PlantaDeleteController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.useCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
