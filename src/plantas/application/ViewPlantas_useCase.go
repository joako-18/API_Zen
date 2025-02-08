package application

import (
	"api/src/plantas/domain/entities"
	"api/src/plantas/domain/repositories"
)

type PlantaViewUseCase struct {
	repo *repositories.PlantaRepository
}

func NewPlantaViewUseCase(repo *repositories.PlantaRepository) *PlantaViewUseCase {
	return &PlantaViewUseCase{repo: repo}
}

func (s *PlantaViewUseCase) GetAll() ([]entities.Planta, error) {
	return s.repo.GetAll()
}

type PlantaNewRecordsUseCase struct {
	repo *repositories.PlantaRepository
}

func NewPlantaNewRecordsUseCase(repo *repositories.PlantaRepository) *PlantaNewRecordsUseCase {
	return &PlantaNewRecordsUseCase{repo: repo}
}

func (s *PlantaNewRecordsUseCase) GetNewRecords(lastUpdated string) ([]entities.Planta, error) {
	return s.repo.GetNewRecords(lastUpdated)
}
