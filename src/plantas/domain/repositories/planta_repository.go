package repositories

import (
	"api/src/plantas/domain/entities"
	"database/sql"
)

type PlantaRepository struct {
	db *sql.DB
}

func NewPlantaRepository(db *sql.DB) *PlantaRepository {
	return &PlantaRepository{db: db}
}

func (r *PlantaRepository) GetAll() ([]entities.Planta, error) {
	rows, err := r.db.Query("SELECT id, nombre, tipo, frecuencia_riego FROM plantas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plantas []entities.Planta
	for rows.Next() {
		var p entities.Planta
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Tipo, &p.FrecuenciaRiego); err != nil {
			return nil, err
		}
		plantas = append(plantas, p)
	}
	return plantas, nil
}

func (r *PlantaRepository) Create(planta entities.Planta) error {
	_, err := r.db.Exec("INSERT INTO plantas (nombre, tipo, frecuencia_riego) VALUES (?, ?, ?)", planta.Nombre, planta.Tipo, planta.FrecuenciaRiego)
	return err
}

func (r *PlantaRepository) Update(planta entities.Planta) error {
	_, err := r.db.Exec("UPDATE plantas SET nombre=?, tipo=?, frecuencia_riego=? WHERE id=?", planta.Nombre, planta.Tipo, planta.FrecuenciaRiego, planta.ID)
	return err
}

func (r *PlantaRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM plantas WHERE id=?", id)
	return err
}
