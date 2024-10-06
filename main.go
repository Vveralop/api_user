package main

import (
	"gogin/config"
	"log"

	"gogin/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Para cargar el archivo .env
)

func main() {
    // Cargar el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error cargando el archivo .env: %v", err)
    }

    // Conectar a la base de datos
    config.ConnectDatabase()

    // Crear una nueva instancia de Gin
    router := gin.Default()

    // Registrar las rutas
    routes.UserRoutes(router)

    // Iniciar el servidor en el puerto 8080
    router.Run(":8080")
}
