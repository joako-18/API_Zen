package repository

import (
	"database/sql"
	"fmt"
	"replica/src/domain/models"
)

type PlantaRepository struct {
	db *sql.DB
}

func NewPlantaRepository(db *sql.DB) *PlantaRepository {
	return &PlantaRepository{db: db}
}

func (r *PlantaRepository) StorePlants(plantas []models.Planta) error {
	for _, p := range plantas {
		_, err := r.db.Exec(
			`INSERT INTO plantas (id, nombre, tipo, riego) 
			 VALUES (?, ?, ?, ?) 
			 ON DUPLICATE KEY UPDATE 
			 nombre=?, tipo=?, riego=?`,
			p.ID, p.Nombre, p.Tipo, p.Riego,
			p.Nombre, p.Tipo, p.Riego,
		)
		if err != nil {
			return fmt.Errorf("error al insertar/actualizar planta ID %d: %v", p.ID, err)
		}
	}
	return nil
}
