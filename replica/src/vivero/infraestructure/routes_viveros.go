package infrastructure

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupViverosRoutes(r *gin.Engine, db *sql.DB) {

	deps := SetupViveroDependencies(db)

	r.GET("/viveros/poll", deps.ShortPollingController.Poll)
	r.GET("/viveros/longpoll", deps.LongPollingController.LongPoll)

}
