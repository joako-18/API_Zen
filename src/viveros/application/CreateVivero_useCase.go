package application

import (
	"api/src/viveros/domain/entities"
	"api/src/viveros/domain/repository"
)

type CreateViveroUseCase struct {
	repo *repository.ViveroRepository
}

func NewCreateViveroUseCase(repo *repository.ViveroRepository) *CreateViveroUseCase {
	return &CreateViveroUseCase{repo: repo}
}

func (s *CreateViveroUseCase) Create(vivero entities.Vivero) error {
	return s.repo.Create(vivero)
}
