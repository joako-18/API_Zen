package entities

type Planta struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Tipo            string `json:"tipo"`
	FrecuenciaRiego string `json:"frecuencia_riego"`
}
