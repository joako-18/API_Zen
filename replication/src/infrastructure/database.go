package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "xk27:mando18D@tcp(localhost:3306)/ZEN_replica")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	return db
}
