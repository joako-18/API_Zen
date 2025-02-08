package usecases

import (
	"replica/src/domain/models"
	"replica/src/domain/repository"
	"time"
)

type ReplicatePlantsUseCase struct {
	repo *repository.PlantaRepository
}

func NewReplicatePlantsUseCase(repo *repository.PlantaRepository) *ReplicatePlantsUseCase {
	return &ReplicatePlantsUseCase{repo: repo}
}

var LastUpdated string = "2024-01-01 00:00:00" // Fecha inicial

func (uc *ReplicatePlantsUseCase) Replicate(lastUpdated string, fetchFunc func(string) ([]models.Planta, error)) error {
	// Obtener las nuevas plantas
	plantas, err := fetchFunc(LastUpdated)
	if err != nil {
		return err
	}

	// Almacenar las plantas en la base de datos
	if len(plantas) > 0 {
		if err := uc.repo.StorePlants(plantas); err != nil {
			return err
		}
		LastUpdated = time.Now().Format("2006-01-02 15:04:05")
	}

	return nil
}
