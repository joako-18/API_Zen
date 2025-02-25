package infrastructure

import (
	"database/sql"
	"replica/src/vivero/application"
	"replica/src/vivero/domain/repository"
	controllers "replica/src/vivero/infraestructure/controller"
)

type ViveroDependencies struct {
	ViveroRepo             *repository.ViveroRepository
	ShortPollingController *controllers.ShortPollViverosController
	LongPollingController  *controllers.LongPollViverosController
	ShortPollingUseCase    *application.ShortPollViverosUseCase
	LongPollingUseCase     *application.LongPollViverosUseCase
}

func SetupViveroDependencies(db *sql.DB) *ViveroDependencies {

	viveroRepo := repository.NewViveroRepository(db)

	shortPollingUseCase := application.NewPollViverosUseCase(viveroRepo)
	longPollingUseCase := application.NewLongPollViverosUseCase(viveroRepo)

	shortPollingController := controllers.NewPollViverosController(shortPollingUseCase)
	longPollingController := controllers.NewLongPollViverosController(longPollingUseCase)

	return &ViveroDependencies{
		ViveroRepo:             viveroRepo,
		ShortPollingUseCase:    shortPollingUseCase,
		LongPollingUseCase:     longPollingUseCase,
		ShortPollingController: shortPollingController,
		LongPollingController:  longPollingController,
	}
}
