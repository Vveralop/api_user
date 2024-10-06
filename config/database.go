package config

import (
	"fmt"
	"gogin/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Declarar la variable DB que contendrá la conexión de la base de datos
var DB *gorm.DB

// Conectar a la base de datos
func ConnectDatabase() {
    var err error

    // Configurar los datos de conexión a la base de datos
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),     // Nombre de usuario
        os.Getenv("DB_PASSWORD"), // Contraseña
        os.Getenv("DB_HOST"),     // Dirección del servidor MySQL
        os.Getenv("DB_PORT"),     // Puerto del servidor MySQL
        os.Getenv("DB_NAME"),     // Nombre de la base de datos
    )

    // Conectar a la base de datos MySQL usando GORM
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Habilitar logs para depuración
    })

    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }

   // Migrar los modelos a la base de datos
    DB.AutoMigrate(&models.User{}) // Aquí puedes agregar otros modelos
    fmt.Println("Conexión a la base de datos y migración exitosa.")

    fmt.Println("Conexión a la base de datos exitosa.")
}
