package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupPlantasRoutes(r *gin.Engine, db *sql.DB) {
	deps := SetupPlantaDependencies(db)

	r.GET("/plantas", deps.PlantaViewController.GetAll)
	r.POST("/plantas", deps.PlantaCreateController.Create)
	r.PUT("/plantas", deps.PlantaUpdateController.Update)
	r.DELETE("/plantas/:id", deps.PlantaDeleteController.Delete)
	r.GET("/plantas/nuevas", deps.PlantaViewController.GetNewPlants)

}
