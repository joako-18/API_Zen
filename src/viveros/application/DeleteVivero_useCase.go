package application

import "api/src/viveros/domain/repository"

type DeleteViveroUseCase struct {
	repo *repository.ViveroRepository
}

func NewDeleteViveroUseCase(repo *repository.ViveroRepository) *DeleteViveroUseCase {
	return &DeleteViveroUseCase{repo: repo}
}

func (s *DeleteViveroUseCase) Delete(id int) error {
	return s.repo.Delete(id)
}
