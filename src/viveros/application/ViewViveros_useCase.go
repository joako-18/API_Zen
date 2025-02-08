package application

import (
	"api/src/viveros/domain/entities"
	"api/src/viveros/domain/repository"
)

type ViewViverosUseCase struct {
	repo *repository.ViveroRepository
}

func NewViewViverosUseCase(repo *repository.ViveroRepository) *ViewViverosUseCase {
	return &ViewViverosUseCase{repo: repo}
}

func (s *ViewViverosUseCase) GetAll() ([]entities.Vivero, error) {
	return s.repo.GetAll()
}
