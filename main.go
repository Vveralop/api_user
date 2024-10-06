package main

import (
	"gogin/config"
	"log"
	"time"

	"gogin/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Para cargar el archivo .env
)

// @title User API
// @version 1.0
// @description This is a sample server for managing Users.
// @host localhost:8080
// @BasePath /
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

    // Configuración de CORS
    corsConfig := cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Cambia esto a tus dominios permitidos
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true, // Permitir credenciales si es necesario
        MaxAge:           12 * time.Hour, // Opcional: define el tiempo que el navegador debe cachear la respuesta CORS
    }
    
    // Usar el middleware CORS
    router.Use(cors.New(corsConfig))

    // Registrar las rutas
    routes.UserRoutes(router)

    // Iniciar el servidor en el puerto 8080
    router.Run(":8080")
}
