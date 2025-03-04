package application

import (
	"api/src/plantas/domain/repositories"
	"fmt"
	"log"
	"time"
)

type ShortPollingUseCase struct {
	repo *repositories.PlantaRepository
}

func NewShortPollingUseCase(repo *repositories.PlantaRepository) *ShortPollingUseCase {
	return &ShortPollingUseCase{repo: repo}
}

func (s *ShortPollingUseCase) GetNewPlantasCount() (int, error) {
	count, err := s.repo.GetNewPlantasCount()
	if err != nil {
		return 0, fmt.Errorf("error al obtener el conteo de plantas: %w", err)
	}
	return count, nil
}

func (s *ShortPollingUseCase) StartPolling(stop chan bool) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			count, err := s.GetNewPlantasCount()
			if err != nil {

				log.Printf("Error al obtener el conteo de plantas: %v", err)
				continue
			}

			log.Printf("Plantas creadas recientemente: %d", count)

		case <-stop:

			log.Println("El polling se ha detenido.")
			return
		}
	}
}
