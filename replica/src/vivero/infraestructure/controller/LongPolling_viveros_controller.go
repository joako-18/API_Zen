package controllers

import (
	"net/http"
	"replica/src/vivero/application"
	"replica/src/vivero/domain/entities"
	"time"

	"github.com/gin-gonic/gin"
)

type LongPollViverosController struct {
	useCase *application.LongPollViverosUseCase
}

func NewLongPollViverosController(useCase *application.LongPollViverosUseCase) *LongPollViverosController {
	return &LongPollViverosController{useCase: useCase}
}

func (h *LongPollViverosController) LongPoll(c *gin.Context) {
	var lastData []entities.Vivero

	timeout := time.After(30 * time.Second) // Máximo tiempo de espera
	ticker := time.Tick(2 * time.Second)    // Consultar cada 2 segundos

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusNoContent, gin.H{"message": "No new data"})
			return
		case <-ticker:
			viveros, err := h.useCase.LongPoll(lastData) // ✅ Pasar lastData como argumento
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if len(viveros) != len(lastData) {
				lastData = viveros // ✅ Actualizar lastData
				c.JSON(http.StatusOK, viveros)
				return
			}
		}
	}
}
