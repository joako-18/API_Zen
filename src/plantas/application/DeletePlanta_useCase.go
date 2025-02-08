package application

import "api/src/plantas/domain/repositories"

type PlantaDeleteUseCase struct {
	repo *repositories.PlantaRepository
}

// Constructor para PlantaDeleteUseCase
func NewPlantaDeleteUseCase(repo *repositories.PlantaRepository) *PlantaDeleteUseCase {
	return &PlantaDeleteUseCase{repo: repo}
}

func (s *PlantaDeleteUseCase) Delete(id int) error {
	return s.repo.Delete(id)
}
