package models

type Planta struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Tipo   string `json:"tipo"`
	Riego  int    `json:"riego"`
}
