package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"backend-go/config"
	"backend-go/database"
	"backend-go/handlers"
	"backend-go/middleware"
	"backend-go/models"
)

var validate *validator.Validate

func GetClientIP(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		return strings.Split(forwardedFor, ",")[0]
	}
	return r.RemoteAddr
}

func myHandler(c *gin.Context) {
	clientIP := GetClientIP(c.Request)
	ginClientIP := c.ClientIP() // Esta debería ser la IP real determinada por Gin

	c.JSON(http.StatusOK, gin.H{
		"clientIP":    clientIP,
		"ginClientIP": ginClientIP,
		"message":     "Hello from handler",
	})
}

func loadConfiguration() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadConfig()
}

func main() {
	loadConfiguration()

	validate = validator.New()

	r := gin.Default()

	// Configurar la confianza de proxies (reemplaza con las IPs REALES de tus proxies)
	r.SetTrustedProxies([]string{
		"192.168.1.45", // IP local de tu servidor
		"192.168.0.1",  // IP de tu proxy o balanceador de carga (EJEMPLO)
		"203.0.113.0",  // IP de tu proxy de infraestructura en la nube (EJEMPLO)
		// Agrega aquí más IPs de proxies confiables si los tienes
	})

	// Inicializar la base de datos
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}

	// **IMPORTANTE: Comentar o eliminar AutoMigrate en producción**
	if os.Getenv("ENV") != "production" {
		database.DB.AutoMigrate(&models.User{}, &models.Article{})
	}

	// Rutas públicas
	r.POST("/login", handlers.Login) // Usamos Login sin pasar `ownerUser`
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Bienvenido a la API"})
	})
	r.GET("/hello", myHandler)

	// Rutas protegidas (solo para el propietario)
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware()) // Usamos el middleware de autenticación JWT
	{
		authGroup.POST("/articles", handlers.CreateArticle)
		authGroup.DELETE("/articles/:id", handlers.DeleteArticle)
	}

	r.Run(config.ServerPort)
}
