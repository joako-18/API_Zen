package application

import (
	"api/src/plantas/domain/repositories"
	"fmt"
	"log"
	"time"
)

type PlantaWateringUseCase struct {
	repo *repositories.PlantaRepository
}

func NewPlantaWateringUseCase(repo *repositories.PlantaRepository) *PlantaWateringUseCase {
	return &PlantaWateringUseCase{repo: repo}
}

func (s *PlantaWateringUseCase) Water(id int) error {

	planta, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	riegoEnMinutos := planta.Riego * 60

	time.Sleep(time.Duration(riegoEnMinutos) * time.Second)

	mensaje := fmt.Sprintf("Es hora de regar la planta %s. Debes regarla cada %d minutos.", planta.Nombre, riegoEnMinutos)
	log.Println(mensaje)

	return nil
}
