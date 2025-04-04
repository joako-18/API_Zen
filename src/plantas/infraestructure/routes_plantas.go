package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupPlantasRoutes(r *gin.Engine, db *sql.DB, secretKey string) {

	deps := SetupPlantaDependencies(db, secretKey)
	r.GET("/plantas", deps.PlantaViewController.GetAll)
	r.POST("/plantas", deps.PlantaCreateController.Create)
	r.PUT("/plantas", deps.PlantaUpdateController.Update)
	r.DELETE("/plantas/:id", deps.PlantaDeleteController.Delete)
}
