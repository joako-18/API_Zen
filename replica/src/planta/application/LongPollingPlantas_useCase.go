package application

import (
	"replica/src/planta/domain/entities"
	"replica/src/planta/domain/repositories"
	"sync"
	"time"
)

type PlantaLongPollingUseCase struct {
	repo       *repositories.PlantaRepository
	lastUpdate time.Time
	mu         sync.Mutex
}

func NewPlantaLongPollingUseCase(repo *repositories.PlantaRepository) *PlantaLongPollingUseCase {
	return &PlantaLongPollingUseCase{repo: repo}
}

func (s *PlantaLongPollingUseCase) WaitForChanges() ([]entities.Planta, error) {
	for {
		time.Sleep(2 * time.Second) // Evita sobrecarga en el CPU
		plantas, lastUpdate, err := s.repo.GetLatestUpdate()
		if err != nil {
			return nil, err
		}

		s.mu.Lock()
		if lastUpdate.After(s.lastUpdate) {
			s.lastUpdate = lastUpdate
			s.mu.Unlock()
			return plantas, nil
		}
		s.mu.Unlock()
	}
}
