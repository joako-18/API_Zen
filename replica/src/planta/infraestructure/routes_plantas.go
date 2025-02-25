package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupPlantasRoutes(r *gin.Engine, db *sql.DB) {
	deps := SetupPlantaDependencies(db)

	r.GET("/plantas/shortpolling", deps.PlantaShortPollingController.GetAll)
	r.GET("/plantas/longpolling", deps.PlantaLongPollingController.LongPolling)

}
