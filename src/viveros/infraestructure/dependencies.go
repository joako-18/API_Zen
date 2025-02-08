package infrastructure

import (
	"api/src/viveros/application"
	"api/src/viveros/domain/repository"
	"api/src/viveros/infraestructure/controllers"
	"database/sql"
)

type ViveroDependencies struct {
	ViveroRepo                 *repository.ViveroRepository
	CreateViveroUseCase        *application.CreateViveroUseCase
	UpdateViveroUseCase        *application.UpdateViveroUseCase
	DeleteViveroUseCase        *application.DeleteViveroUseCase
	ViewViverosUseCase         *application.ViewViverosUseCase
	ReplicateViverosUseCase    *application.ReplicateViverosUseCase
	CreateViveroController     *controllers.CreateViveroController
	UpdateViveroController     *controllers.UpdateViveroController
	DeleteViveroController     *controllers.DeleteViveroController
	ViewViverosController      *controllers.ViewViverosController
	ReplicateViverosController *controllers.ReplicateViverosController
}

func SetupViveroDependencies(db *sql.DB) *ViveroDependencies {

	viveroRepo := repository.NewViveroRepository(db)

	createViveroUseCase := application.NewCreateViveroUseCase(viveroRepo)
	updateViveroUseCase := application.NewUpdateViveroUseCase(viveroRepo)
	deleteViveroUseCase := application.NewDeleteViveroUseCase(viveroRepo)
	viewViverosUseCase := application.NewViewViverosUseCase(viveroRepo)
	replicateViverosUseCase := application.NewReplicateViverosUseCase(viveroRepo)

	createViveroController := controllers.NewCreateViveroController(createViveroUseCase)
	updateViveroController := controllers.NewUpdateViveroController(updateViveroUseCase)
	deleteViveroController := controllers.NewDeleteViveroController(deleteViveroUseCase)
	viewViverosController := controllers.NewViewViverosController(viewViverosUseCase)
	replicateViverosController := controllers.NewReplicateViverosController(replicateViverosUseCase)

	return &ViveroDependencies{
		ViveroRepo:                 viveroRepo,
		CreateViveroUseCase:        createViveroUseCase,
		UpdateViveroUseCase:        updateViveroUseCase,
		DeleteViveroUseCase:        deleteViveroUseCase,
		ViewViverosUseCase:         viewViverosUseCase,
		ReplicateViverosUseCase:    replicateViverosUseCase,
		CreateViveroController:     createViveroController,
		UpdateViveroController:     updateViveroController,
		DeleteViveroController:     deleteViveroController,
		ViewViverosController:      viewViverosController,
		ReplicateViverosController: replicateViverosController,
	}
}
