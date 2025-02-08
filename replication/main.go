package main

import (
	"log"
	"replica/src/application/usecases"
	"replica/src/domain/repository"
	"replica/src/infrastructure"
	"replica/src/infrastructure/services"
	"time"
)

func main() {
	// Conectar a la base de datos
	db := infrastructure.ConnectDatabase()
	defer db.Close()

	// Configurar dependencias para plantas
	plantaRepo := repository.NewPlantaRepository(db)
	replicatePlants := usecases.NewReplicatePlantsUseCase(plantaRepo)

	// Configurar dependencias para viveros
	viveroRepo := repository.NewViveroRepository(db)
	replicateViveros := usecases.NewReplicateViverosUseCase(viveroRepo)

	// Iniciar replicación periódica
	log.Println("Iniciando replicación de plantas y viveros...")

	go func() {
		for {
			err := replicatePlants.Replicate(usecases.LastUpdated, services.FetchNewPlants) // Pasar la fecha correctamente
			if err != nil {
				log.Println("Error durante la replicación de plantas:", err)
			} else {
				log.Println("Replicación de plantas completada exitosamente.")
			}
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := replicateViveros.Replicate(usecases.LastUpdatedViveros, services.FetchNewViveros) // Pasar la fecha correctamente
			if err != nil {
				log.Println("Error durante la replicación de viveros:", err)
			} else {
				log.Println("Replicación de viveros completada exitosamente.")
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
