package infrastructure

import (
	"api/src/plantas/application"
	"api/src/plantas/domain/repositories"
	"api/src/plantas/infraestructure/controller"
	services "api/src/plantas/infraestructure/service"
	"database/sql"
	"log"
)

type PlantaDependencies struct {
	PlantaRepo             repositories.PlantaRepository
	PlantaCreateUseCase    *application.PlantaCreateUseCase
	PlantaUpdateUseCase    *application.PlantaUpdateUseCase
	PlantaDeleteUseCase    *application.PlantaDeleteUseCase
	PlantaViewUseCase      *application.PlantaViewUseCase
	PlantaCreateController *controller.PlantaCreateController
	PlantaUpdateController *controller.PlantaUpdateController
	PlantaDeleteController *controller.PlantaDeleteController
	PlantaViewController   *controller.PlantaViewController
}

func SetupPlantaDependencies(db *sql.DB, secretKey string) *PlantaDependencies {
	plantaRepo := repositories.NewPlantaRepository(db)
	encryptService, err := services.NewEncryptService(secretKey)
	if err != nil {
		log.Fatalf("Error initializing EncryptService: %v", err)
	}

	plantaCreateUseCase := application.NewPlantaCreateUseCase(plantaRepo, encryptService)
	plantaUpdateUseCase := application.NewPlantaUpdateUseCase(plantaRepo)
	plantaDeleteUseCase := application.NewPlantaDeleteUseCase(plantaRepo)
	plantaViewUseCase := application.NewPlantaViewUseCase(plantaRepo)

	plantaCreateController := controller.NewPlantaCreateController(plantaCreateUseCase)
	plantaUpdateController := controller.NewPlantaUpdateController(plantaUpdateUseCase)
	plantaDeleteController := controller.NewPlantaDeleteController(plantaDeleteUseCase)
	plantaViewController := controller.NewPlantaViewController(plantaViewUseCase)

	return &PlantaDependencies{
		PlantaRepo:             plantaRepo,
		PlantaCreateUseCase:    plantaCreateUseCase,
		PlantaUpdateUseCase:    plantaUpdateUseCase,
		PlantaDeleteUseCase:    plantaDeleteUseCase,
		PlantaViewUseCase:      plantaViewUseCase,
		PlantaCreateController: plantaCreateController,
		PlantaUpdateController: plantaUpdateController,
		PlantaDeleteController: plantaDeleteController,
		PlantaViewController:   plantaViewController,
	}
}
