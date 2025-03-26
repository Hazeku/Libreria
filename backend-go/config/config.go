package config

import "os"

var (
    DatabaseURL = os.Getenv("DATABASE_URL")
    JWTSecret   = os.Getenv("JWT_SECRET")
    ServerPort  = os.Getenv("SERVER_PORT")
)

func LoadConfig() {
    if DatabaseURL == "" {
        DatabaseURL = "host=localhost ..." // Valores por defecto para desarrollo
    }
    if JWTSecret == "" {
        JWTSecret = "libreria_master_2025"
    }
    if ServerPort == "" {
        ServerPort = ":8000"
    }
}