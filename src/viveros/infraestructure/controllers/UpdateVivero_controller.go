package controllers

import (
	"api/src/viveros/application"
	"api/src/viveros/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateViveroController struct {
	useCase *application.UpdateViveroUseCase
}

func NewUpdateViveroController(useCase *application.UpdateViveroUseCase) *UpdateViveroController {
	return &UpdateViveroController{useCase: useCase}
}

func (h *UpdateViveroController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var vivero entities.Vivero
	if err := c.ShouldBindJSON(&vivero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vivero.ID = id
	if err := h.useCase.Update(vivero); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vivero"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vivero updated successfully"})
}
