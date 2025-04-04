package application

import (
	"api/src/plantas/domain/entities"
	"api/src/plantas/domain/repositories"
	services "api/src/plantas/infraestructure/service"
)

type PlantaCreateUseCase struct {
	repo           repositories.PlantaRepository
	encryptService *services.EncryptService
}

func NewPlantaCreateUseCase(repo repositories.PlantaRepository, encryptService *services.EncryptService) *PlantaCreateUseCase {
	return &PlantaCreateUseCase{
		repo:           repo,
		encryptService: encryptService,
	}
}

func (s *PlantaCreateUseCase) Create(planta entities.Planta) error {
	hashedTipo, err := s.encryptService.Encrypt(planta.Tipo)
	if err != nil {
		return err
	}
	planta.Tipo = hashedTipo
	return s.repo.Create(planta)
}
