package repository

import (
	"api/src/viveros/domain/entities"
	"database/sql"
)

type ViveroRepository struct {
	db *sql.DB
}

func NewViveroRepository(db *sql.DB) *ViveroRepository {
	return &ViveroRepository{db: db}
}

func (r *ViveroRepository) GetAll() ([]entities.Vivero, error) {
	rows, err := r.db.Query("SELECT id, nombre, descripcion, direccion FROM viveros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var viveros []entities.Vivero
	for rows.Next() {
		var v entities.Vivero
		if err := rows.Scan(&v.ID, &v.Nombre, &v.Descripcion, &v.Direccion); err != nil {
			return nil, err
		}
		viveros = append(viveros, v)
	}
	return viveros, nil
}

func (r *ViveroRepository) Create(vivero entities.Vivero) error {
	_, err := r.db.Exec("INSERT INTO viveros (nombre, descripcion, direccion) VALUES (?, ?, ?)", vivero.Nombre, vivero.Descripcion, vivero.Direccion)
	return err
}

func (r *ViveroRepository) Update(vivero entities.Vivero) error {
	_, err := r.db.Exec("UPDATE viveros SET nombre=?, descripcion=?, direccion=? WHERE id=?", vivero.Nombre, vivero.Descripcion, vivero.Direccion, vivero.ID)
	return err
}

func (r *ViveroRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM viveros WHERE id=?", id)
	return err
}

// Función para replicar viveros (inserta o actualiza)
func (r *ViveroRepository) ReplicateViveros(viveros []entities.Vivero) error {
	for _, v := range viveros {
		_, err := r.db.Exec(
			"INSERT INTO viveros (id, nombre, descripcion, direccion) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE nombre=?, descripcion=?, direccion=?",
			v.ID, v.Nombre, v.Descripcion, v.Direccion, v.Nombre, v.Descripcion, v.Direccion,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
