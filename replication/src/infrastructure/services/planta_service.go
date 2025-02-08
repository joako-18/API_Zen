package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"replica/src/domain/models"
)

func FetchNewPlants(lastUpdated string) ([]models.Planta, error) {
	url := fmt.Sprintf("http://servidor-principal:8080/plantas/nuevas?last_updated=%s", lastUpdated)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al obtener nuevas plantas: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("respuesta no exitosa del servidor principal: %d", resp.StatusCode)
	}

	var plantas []models.Planta
	if err := json.NewDecoder(resp.Body).Decode(&plantas); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta JSON: %v", err)
	}

	return plantas, nil
}
