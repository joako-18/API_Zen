package controllers

import (
	"api/src/viveros/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReplicateViverosController struct {
	useCase *application.ReplicateViverosUseCase
}

func NewReplicateViverosController(useCase *application.ReplicateViverosUseCase) *ReplicateViverosController {
	return &ReplicateViverosController{useCase: useCase}
}

func (c *ReplicateViverosController) Replicate(cxt *gin.Context) {
	err := c.useCase.Replicate()
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "Replicación completada con éxito."})
}
