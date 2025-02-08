package application

import (
	"api/src/viveros/domain/entities"
	"api/src/viveros/domain/repository"
)

type UpdateViveroUseCase struct {
	repo *repository.ViveroRepository
}

func NewUpdateViveroUseCase(repo *repository.ViveroRepository) *UpdateViveroUseCase {
	return &UpdateViveroUseCase{repo: repo}
}

func (s *UpdateViveroUseCase) Update(vivero entities.Vivero) error {
	return s.repo.Update(vivero)
}
