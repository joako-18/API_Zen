package repositories

import (
	"api/src/plantas/domain/entities"
	"database/sql"
)

// Definir la interfaz
type PlantaRepository interface {
	GetAll() ([]entities.Planta, error)
	Create(planta entities.Planta) error
	Update(planta entities.Planta) error
	Delete(id int) error
}

// Estructura que implementa la interfaz
type PlantaRepositoryImpl struct {
	db *sql.DB
}

// Constructor para la implementación de la interfaz
func NewPlantaRepository(db *sql.DB) *PlantaRepositoryImpl {
	return &PlantaRepositoryImpl{db: db}
}

// Implementación del método GetAll
func (r *PlantaRepositoryImpl) GetAll() ([]entities.Planta, error) {
	rows, err := r.db.Query("SELECT id, nombre, tipo, riego FROM plantas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plantas []entities.Planta
	for rows.Next() {
		var p entities.Planta
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Tipo, &p.Riego); err != nil {
			return nil, err
		}
		plantas = append(plantas, p)
	}
	return plantas, nil
}

// Implementación del método Create
func (r *PlantaRepositoryImpl) Create(planta entities.Planta) error {
	_, err := r.db.Exec("INSERT INTO plantas (nombre, tipo, riego) VALUES (?, ?, ?)", planta.Nombre, planta.Tipo, planta.Riego)
	return err
}

// Implementación del método Update
func (r *PlantaRepositoryImpl) Update(planta entities.Planta) error {
	_, err := r.db.Exec("UPDATE plantas SET nombre=?, tipo=?, riego=? WHERE id=?", planta.Nombre, planta.Tipo, planta.Riego, planta.ID)
	return err
}

// Implementación del método Delete
func (r *PlantaRepositoryImpl) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM plantas WHERE id=?", id)
	return err
}
