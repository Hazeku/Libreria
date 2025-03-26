package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "backend-go/models"
)

// InitTestDB inicializa una base de datos SQLite en memoria para pruebas.
func InitTestDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrar los esquemas necesarios para las pruebas
    db.AutoMigrate(&models.User{}, &models.Article{}) // Asegúrate de migrar todos los modelos necesarios

    return db, nil
}

// CloseTestDB cierra la conexión a la base de datos de prueba (no necesario para SQLite en memoria, pero se incluye por consistencia).
func CloseTestDB(db *gorm.DB) error {
    sqlDB, err := db.DB()
    if err != nil {
        return err
    }
    return sqlDB.Close()
}