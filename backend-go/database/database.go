package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"

	"backend-go/models"
)

var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos y realiza las migraciones
func InitDB() error {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️ Advertencia: No se pudo cargar .env, usando valores por defecto.")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		fmt.Println("⚠️ DATABASE_URL no está configurada. Usando configuración por defecto.")
		dbURL = "postgres://usuario:contraseña@host:puerto/dbname?sslmode=disable&pgbouncer=true&prepared_statements=false&prefer_simple_protocol=true"
	} else {
		// Agregar configuración segura (sin statement_cache_mode)
		if !strings.Contains(dbURL, "?") {
			dbURL += "?pool_max_conns=10&statement_cache_mode=disable"
		} else {
			dbURL += "&pool_max_conns=10&statement_cache_mode=disable"
		}
	}

	fmt.Println("🔌 Conectando a la base de datos:", dbURL)

	// Conexión con GORM
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("❌ Error al conectar con la base de datos: %w", err)
	}

	DB = db

	// Configuración de conexiones
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("❌ Error al obtener sql.DB: %w", err)
	}

	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("❌ No se pudo establecer conexión: %w", err)
	}

	fmt.Println("🚀 ¡Conexión a la base de datos exitosa!")
	return nil
}

// Reiniciar la conexión a la base de datos
func RestartDBConnection() error {
	// Cerrar la conexión actual
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("❌ Error al obtener sql.DB para cierre: %w", err)
	}

	// Cerrar la conexión existente
	sqlDB.Close()
	fmt.Println("🔌 Conexión a la base de datos cerrada.")

	// Esperar unos segundos antes de reabrir la conexión
	time.Sleep(2 * time.Second)

	// Volver a inicializar la conexión
	err = InitDB()
	if err != nil {
		return fmt.Errorf("❌ Error al reiniciar la conexión a la base de datos: %w", err)
	}

	fmt.Println("🔄 Conexión a la base de datos reiniciada exitosamente.")
	return nil
}

// Migraciones con chequeo por tabla
func Migrate(DB *gorm.DB) error {
	// Lista de modelos a migrar
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Article{},
		&models.Order{},
		&models.OrderArticle{},
	}

	// Comprobamos si la tabla existe y realizamos la migración si es necesario
	for _, model := range modelsToMigrate {
		if !DB.Migrator().HasTable(model) {
			fmt.Printf("🛠️ Tabla no encontrada, creando: %T\n", model)
			if err := DB.AutoMigrate(model); err != nil {
				fmt.Printf("❌ Error migrando %T: %v\n", model, err)
				return err
			}
		} else {
			fmt.Printf("✅ Tabla existente, saltando migración: %T\n", model)
		}
	}

	// Devolvemos nil si todo salió bien
	return nil
}
