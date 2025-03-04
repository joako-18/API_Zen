package controller

import (
	"api/src/plantas/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortPollingController struct {
	useCase *application.ShortPollingUseCase
}

func NewShortPollingController(useCase *application.ShortPollingUseCase) *ShortPollingController {
	return &ShortPollingController{useCase: useCase}
}

func (h *ShortPollingController) StartPolling(c *gin.Context) {
	stop := make(chan bool)

	go func() {
		defer close(stop)
		h.useCase.StartPolling(stop)
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Polling iniciado, consulta cada 5 segundos."})
}
