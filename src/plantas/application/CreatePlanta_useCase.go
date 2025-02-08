package application

import (
	"api/src/plantas/domain/entities"
	"api/src/plantas/domain/repositories"
)

type PlantaCreateUseCase struct {
	repo *repositories.PlantaRepository
}

func NewPlantaCreateUseCase(repo *repositories.PlantaRepository) *PlantaCreateUseCase {
	return &PlantaCreateUseCase{repo: repo}
}

func (s *PlantaCreateUseCase) Create(planta entities.Planta) error {
	return s.repo.Create(planta)
}
