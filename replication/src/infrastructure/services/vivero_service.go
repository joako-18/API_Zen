package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"replica/src/domain/models"
)

func FetchNewViveros(lastUpdated string) ([]models.Vivero, error) {
	url := fmt.Sprintf("http://servidor-principal:8080/viveros/nuevos?last_updated=%s", lastUpdated)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al obtener nuevos viveros: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("respuesta no exitosa del servidor principal: %d", resp.StatusCode)
	}

	var viveros []models.Vivero
	if err := json.NewDecoder(resp.Body).Decode(&viveros); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta JSON: %v", err)
	}

	return viveros, nil
}
