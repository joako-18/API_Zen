package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

func ConnectDB() (*sql.DB, error) {
	var err error
	once.Do(func() {
		if loadErr := godotenv.Load(); loadErr != nil {
			log.Println("No se pudo cargar el archivo .env")
		}

		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

		dbInstance, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error conectando a la base de datos: %v", err)
			return
		}

		if pingErr := dbInstance.Ping(); pingErr != nil {
			log.Fatalf("Error verificando la conexión: %v", pingErr)
			return
		}

		log.Println("Base de datos conectada correctamente")
	})

	return dbInstance, err
}

func GetDBInstance() *sql.DB {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("No se pudo obtener la instancia de la base de datos: %v", err)
	}
	return db
}
