package controllers

import (
	"net/http"
	"replica/src/vivero/application"

	"github.com/gin-gonic/gin"
)

type ShortPollViverosController struct {
	useCase *application.ShortPollViverosUseCase
}

func NewPollViverosController(useCase *application.ShortPollViverosUseCase) *ShortPollViverosController {
	return &ShortPollViverosController{useCase: useCase}
}

func (h *ShortPollViverosController) Poll(c *gin.Context) {
	viveros, err := h.useCase.Poll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, viveros)
}
