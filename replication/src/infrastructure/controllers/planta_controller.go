package controllers

import (
	"encoding/json"
	"net/http"
	"replica/src/application/usecases"
	"replica/src/infrastructure/services"
)

type PlantaController struct {
	useCase *usecases.ReplicatePlantsUseCase
}

func NewPlantaController(useCase *usecases.ReplicatePlantsUseCase) *PlantaController {
	return &PlantaController{useCase: useCase}
}

func (pc *PlantaController) ReplicatePlantsHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí podríamos obtener el valor de lastUpdated de algún almacenamiento o configuración
	lastUpdated := "2024-01-01 00:00:00"

	err := pc.useCase.Replicate(lastUpdated, services.FetchNewPlants)
	if err != nil {
		http.Error(w, "Error al replicar plantas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Replicación de plantas completada"})
}
