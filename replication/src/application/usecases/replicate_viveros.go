package usecases

import (
	"replica/src/domain/models"
	"replica/src/domain/repository"
)

type ReplicateViverosUseCase struct {
	repo *repository.ViveroRepository
}

func NewReplicateViverosUseCase(repo *repository.ViveroRepository) *ReplicateViverosUseCase {
	return &ReplicateViverosUseCase{repo: repo}
}

var LastUpdatedViveros string = "2024-01-01 00:00:00" // Fecha inicial

func (uc *ReplicateViverosUseCase) Replicate(lastUpdated string, fetchFunc func(string) ([]models.Vivero, error)) error {
	viveros, err := fetchFunc(lastUpdated) // ✅ Pasamos el parámetro correcto
	if err != nil {
		return err
	}

	if len(viveros) > 0 {
		if err := uc.repo.StoreViveros(viveros); err != nil {
			return err
		}
	}

	return nil
}
