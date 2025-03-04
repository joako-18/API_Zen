package infrastructure

import (
	"api/src/plantas/application"
	"api/src/plantas/domain/repositories"
	"api/src/plantas/infraestructure/controller"
	"database/sql"
)

type PlantaDependencies struct {
	PlantaRepo                   *repositories.PlantaRepository
	PlantaCreateUseCase          *application.PlantaCreateUseCase
	PlantaUpdateUseCase          *application.PlantaUpdateUseCase
	PlantaDeleteUseCase          *application.PlantaDeleteUseCase
	PlantaViewUseCase            *application.PlantaViewUseCase
	PlantaWateringUseCase        *application.PlantaWateringUseCase
	PlantaCreateController       *controller.PlantaCreateController
	PlantaUpdateController       *controller.PlantaUpdateController
	PlantaDeleteController       *controller.PlantaDeleteController
	PlantaViewController         *controller.PlantaViewController
	PlantaWateringController     *controller.PlantaWateringController
	PlantaShortPollingUseCase    *application.ShortPollingUseCase
	PlantaShortPollingController *controller.ShortPollingController
}

func SetupPlantaDependencies(db *sql.DB) *PlantaDependencies {
	plantaRepo := repositories.NewPlantaRepository(db)

	plantaCreateUseCase := application.NewPlantaCreateUseCase(plantaRepo)
	plantaUpdateUseCase := application.NewPlantaUpdateUseCase(plantaRepo)
	plantaDeleteUseCase := application.NewPlantaDeleteUseCase(plantaRepo)
	plantaViewUseCase := application.NewPlantaViewUseCase(plantaRepo)
	plantaWateringUseCase := application.NewPlantaWateringUseCase(plantaRepo)
	plantaShortPollingUseCase := application.NewShortPollingUseCase(plantaRepo)
	plantaShortPollingController := controller.NewShortPollingController(plantaShortPollingUseCase)

	plantaCreateController := controller.NewPlantaCreateController(plantaCreateUseCase)
	plantaUpdateController := controller.NewPlantaUpdateController(plantaUpdateUseCase)
	plantaDeleteController := controller.NewPlantaDeleteController(plantaDeleteUseCase)
	plantaViewController := controller.NewPlantaViewController(plantaViewUseCase)
	plantaWateringController := controller.NewPlantaWateringController(plantaWateringUseCase)

	return &PlantaDependencies{
		PlantaRepo:                   plantaRepo,
		PlantaCreateUseCase:          plantaCreateUseCase,
		PlantaUpdateUseCase:          plantaUpdateUseCase,
		PlantaDeleteUseCase:          plantaDeleteUseCase,
		PlantaViewUseCase:            plantaViewUseCase,
		PlantaWateringUseCase:        plantaWateringUseCase,
		PlantaCreateController:       plantaCreateController,
		PlantaUpdateController:       plantaUpdateController,
		PlantaDeleteController:       plantaDeleteController,
		PlantaViewController:         plantaViewController,
		PlantaWateringController:     plantaWateringController,
		PlantaShortPollingUseCase:    plantaShortPollingUseCase,
		PlantaShortPollingController: plantaShortPollingController,
	}
}
