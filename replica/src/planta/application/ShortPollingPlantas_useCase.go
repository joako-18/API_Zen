package application

import (
	"replica/src/planta/domain/entities"
	"replica/src/planta/domain/repositories"
)

type PlantaShortPollingUseCase struct {
	repo *repositories.PlantaRepository
}

func NewPlantaShortPollingUseCase(repo *repositories.PlantaRepository) *PlantaShortPollingUseCase {
	return &PlantaShortPollingUseCase{repo: repo}
}

func (s *PlantaShortPollingUseCase) GetAll() ([]entities.Planta, error) {
	return s.repo.GetAll()
}
