package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend-go/models" // Importa el paquete models
)

var DB *gorm.DB

func InitDB() error {
	// Intentar cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		// No retornar aquí, ya que la URL podría estar directamente en la variable de entorno
	}

	// Obtener la URL de la base de datos de la variable de entorno DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")

	// Verificar si la variable de entorno está vacía y usar un valor por defecto si es necesario
	if dbURL == "" {
		fmt.Println("DATABASE_URL environment variable is not set. Using a default connection string.")
		dbURL = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable" // Valor por defecto para desarrollo
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return err // Retornar el error
	}

	DB = db

	// AutoMigrate crea las tablas automáticamente basadas en tus modelos
	DB.AutoMigrate(&models.User{}, &models.Article{})

	fmt.Println("Successfully connected to the database!")

	return nil
}
