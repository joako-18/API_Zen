package repository

import (
	"database/sql"
	"fmt"
	"replica/src/domain/models"
)

type ViveroRepository struct {
	db *sql.DB
}

func NewViveroRepository(db *sql.DB) *ViveroRepository {
	return &ViveroRepository{db: db}
}

func (r *ViveroRepository) StoreViveros(viveros []models.Vivero) error {
	for _, v := range viveros {
		_, err := r.db.Exec(
			`INSERT INTO viveros (id, nombre, ubicacion, capacidad) 
			 VALUES (?, ?, ?, ?) 
			 ON DUPLICATE KEY UPDATE 
			 nombre=?, ubicacion=?, capacidad=?`,
			v.ID, v.Nombre, v.Ubicacion, v.Capacidad,
			v.Nombre, v.Ubicacion, v.Capacidad,
		)
		if err != nil {
			return fmt.Errorf("error al insertar/actualizar vivero ID %d: %v", v.ID, err)
		}
	}
	return nil
}
