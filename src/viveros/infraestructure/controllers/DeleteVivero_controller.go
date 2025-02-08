package controllers

import (
	"api/src/viveros/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteViveroController struct {
	useCase *application.DeleteViveroUseCase
}

func NewDeleteViveroController(useCase *application.DeleteViveroUseCase) *DeleteViveroController {
	return &DeleteViveroController{useCase: useCase}
}

func (h *DeleteViveroController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	h.useCase.Delete(id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
