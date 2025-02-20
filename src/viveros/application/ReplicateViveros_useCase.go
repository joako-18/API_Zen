package application

import (
	"api/src/viveros/domain/entities"
	"api/src/viveros/domain/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ReplicateViverosUseCase struct {
	repo        *repository.ViveroRepository
	lastUpdated string
}

func NewReplicateViverosUseCase(repo *repository.ViveroRepository) *ReplicateViverosUseCase {
	return &ReplicateViverosUseCase{
		repo:        repo,
		lastUpdated: "2024-01-01 00:00:00", // Fecha inicial
	}
}

// Fetch nuevos viveros del servidor principal
func (u *ReplicateViverosUseCase) FetchNewViveros() ([]entities.Vivero, error) {
	url := fmt.Sprintf("http://servidor-principal:8080/viveros/nuevos?last_updated=%s", u.lastUpdated)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en el servidor principal: %d", resp.StatusCode)
	}

	var viveros []entities.Vivero
	if err := json.NewDecoder(resp.Body).Decode(&viveros); err != nil {
		return nil, err
	}

	return viveros, nil
}

func (u *ReplicateViverosUseCase) Replicate() error {
	viveros, err := u.FetchNewViveros()
	if err != nil {
		return err
	}

	if len(viveros) > 0 {
		if err := u.repo.ReplicateViveros(viveros); err != nil {
			return err
		}
		u.lastUpdated = time.Now().Format("2006-01-02 15:04:05")
	}
	return nil
}
