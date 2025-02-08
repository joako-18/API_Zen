package controllers

import (
	"encoding/json"
	"net/http"
	"replica/src/application/usecases"
	"replica/src/infrastructure/services"
)

type ViveroController struct {
	useCase *usecases.ReplicateViverosUseCase
}

func NewViveroController(useCase *usecases.ReplicateViverosUseCase) *ViveroController {
	return &ViveroController{useCase: useCase}
}

func (vc *ViveroController) ReplicateViverosHandler(w http.ResponseWriter, r *http.Request) {
	lastUpdated := "2024-01-01 00:00:00" // Se puede obtener dinámicamente más adelante

	err := vc.useCase.Replicate(lastUpdated, services.FetchNewViveros)
	if err != nil {
		http.Error(w, "Error al replicar viveros: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Replicación de viveros completada"})
}
