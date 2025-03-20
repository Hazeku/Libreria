package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Modelo de artículo
type Article struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

// Inicializar la base de datos
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar la base de datos:", err)
	}
	DB.AutoMigrate(&Article{}) // Migrar estructura de artículos
}
