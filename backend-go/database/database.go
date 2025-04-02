package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend-go/models" // Importa el paquete models
)

var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos y realiza las migraciones
func InitDB() error {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Advertencia: No se pudo cargar el archivo .env, usando valores predeterminados.")
	}

	// Obtener la URL de la base de datos
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		fmt.Println("DATABASE_URL no está configurada. Usando configuración por defecto.")
		dbURL = "postgres://usuario:contraseña@host:puerto/dbname?sslmode=disable&pgbouncer=true&prepared_statements=false&prefer_simple_protocol=true"
	} else {
		// Agregar parámetros adicionales a la URL si no están presentes
		if !strings.Contains(dbURL, "?") {
			dbURL += "?pool_max_conns=10&statement_cache_mode=describe"
		} else {
			dbURL += "&pool_max_conns=10&statement_cache_mode=describe"
		}
	}

	fmt.Println("Conectando a la base de datos:", dbURL)

	// Conectar a la base de datos con GORM
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	DB = db // Asignar la base de datos a la variable global

	// Configuración de la conexión
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error al obtener sql.DB: %w", err)
	}

	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxOpenConns(10) // Número máximo de conexiones activas
	sqlDB.SetMaxIdleConns(5)  // Número máximo de conexiones inactivas

	// Verificar la conexión
	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("no se pudo establecer conexión con la base de datos: %w", err)
	}

	// Verificar si la tabla 'users' ya existe
	if DB.Migrator().HasTable(&models.User{}) {
		fmt.Println("✅ La tabla 'users' ya existe, saltando migración.")
	} else {
		// Si la tabla no existe, crear las tablas de los modelos
		fmt.Println("🛠️ La tabla 'users' no existe, creando...")
		if err := DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Order{}, &models.OrderArticle{}); err != nil {
			fmt.Printf("❌ Error en la migración de modelos: %v\n", err)
			return err
		}
		fmt.Println("✅ Migración completada con éxito.")
	}

	fmt.Println("🚀 ¡Conexión a la base de datos exitosa!")
	return nil
}
