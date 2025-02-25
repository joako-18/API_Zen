package application

import (
	"replica/src/vivero/domain/entities"
	"replica/src/vivero/domain/repository"
)

type ShortPollViverosUseCase struct {
	repo *repository.ViveroRepository
}

func NewPollViverosUseCase(repo *repository.ViveroRepository) *ShortPollViverosUseCase {
	return &ShortPollViverosUseCase{repo: repo}
}

func (s *ShortPollViverosUseCase) Poll() ([]entities.Vivero, error) {
	return s.repo.GetAll()
}
