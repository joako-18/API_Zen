package application

import (
	"api/src/plantas/domain/entities"
	"api/src/plantas/domain/repositories"
)

type PlantaUpdateUseCase struct {
	repo *repositories.PlantaRepository
}

func NewPlantaUpdateUseCase(repo *repositories.PlantaRepository) *PlantaUpdateUseCase {
	return &PlantaUpdateUseCase{repo: repo}
}

func (s *PlantaUpdateUseCase) Update(planta entities.Planta) error {
	return s.repo.Update(planta)
}
