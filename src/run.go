package src

import (
	"log"
	"os"

	core "api/core"
	plantasInfra "api/src/plantas/infraestructure"
	viverosInfra "api/src/viveros/infraestructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando valores predeterminados")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "default_secret_key"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	db := core.GetDBInstance()

	plantasInfra.SetupPlantasRoutes(r, db, secretKey)
	viverosInfra.SetupViverosRoutes(r, db)

	port := ":8080"
	log.Println("Servidor corriendo en el puerto", port)
	r.Run(port)
}
