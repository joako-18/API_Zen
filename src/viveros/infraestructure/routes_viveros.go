package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupViverosRoutes(r *gin.Engine, db *sql.DB) {

	deps := SetupViveroDependencies(db)

	r.GET("/viveros", deps.ViewViverosController.GetAll)
	r.POST("/viveros", deps.CreateViveroController.Create)
	r.PUT("/viveros/:id", deps.UpdateViveroController.Update)
	r.DELETE("/viveros/:id", deps.DeleteViveroController.Delete)
}
