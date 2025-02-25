package infrastructure

import (
	"database/sql"
	"replica/src/planta/application"
	"replica/src/planta/domain/repositories"
	"replica/src/planta/infraestructure/controller"
)

type PlantaDependencies struct {
	PlantaRepo                   *repositories.PlantaRepository
	PlantaShortPollingUseCase    *application.PlantaShortPollingUseCase
	PlantaLongPollingUseCase     *application.PlantaLongPollingUseCase
	PlantaShortPollingController *controller.PlantaShortPollingController
	PlantaLongPollingController  *controller.PlantaLongPollingController
}

func SetupPlantaDependencies(db *sql.DB) *PlantaDependencies {
	plantaRepo := repositories.NewPlantaRepository(db)

	plantaShortPollingUseCase := application.NewPlantaShortPollingUseCase(plantaRepo)
	plantaLongPollingUseCase := application.NewPlantaLongPollingUseCase(plantaRepo)

	plantaShortPollingController := controller.NewPlantaShortPollingController(plantaShortPollingUseCase)
	plantaLongPollingController := controller.NewPlantaLongPollingController(plantaLongPollingUseCase)

	return &PlantaDependencies{

		PlantaRepo:                   plantaRepo,
		PlantaShortPollingUseCase:    plantaShortPollingUseCase,
		PlantaLongPollingUseCase:     plantaLongPollingUseCase,
		PlantaShortPollingController: plantaShortPollingController,
		PlantaLongPollingController:  plantaLongPollingController,
	}
}
