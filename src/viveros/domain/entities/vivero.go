package entities

type Vivero struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Direccion   string `json:"direccion"`
}
