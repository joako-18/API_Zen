package controllers

import (
	"api/src/viveros/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewViverosController struct {
	useCase *application.ViewViverosUseCase
}

func NewViewViverosController(useCase *application.ViewViverosUseCase) *ViewViverosController {
	return &ViewViverosController{useCase: useCase}
}

func (h *ViewViverosController) GetAll(c *gin.Context) {
	viveros, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve viveros"})
		return
	}

	c.JSON(http.StatusOK, viveros)
}
