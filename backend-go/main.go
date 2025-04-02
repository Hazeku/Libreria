package main

import (
	"fmt"
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
	ginClientIP := c.ClientIP() // Esta debería ser la IP real determinada por Gin

	c.JSON(http.StatusOK, gin.H{
		"ginClientIP": ginClientIP,
		"message":     "Hello from handler",
	})
}

func loadConfiguration() {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtener JWT_SECRET desde las variables de entorno
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in the .env file")
	}

	fmt.Println("JWT_SECRET: ", jwtSecret)

	config.LoadConfig() // Asegúrate de que `config.LoadConfig` también esté usando el JWT_SECRET si es necesario.
}

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error al cargar el archivo .env")
	}

	// Imprimir el valor de la variable PORT
	port := os.Getenv("PORT")
	fmt.Println("Valor de la variable PORT:", port)

	// Inicializar Gin
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

	// Importar los artículos después de inicializar la DB (solo en desarrollo)
	if os.Getenv("ENV") != "production" {
		err := database.ImportArticles()
		if err != nil {
			log.Printf("⚠️  No se pudieron importar los artículos: %v", err)
		}
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
		authGroup.GET("/articles", handlers.GetArticles)
		authGroup.POST("/articles", handlers.CreateArticle)
		authGroup.DELETE("/articles/:id", handlers.DeleteArticle)
		authGroup.PUT("/articles/:id/assign", handlers.AssignArticleToUser)
	}

	// Obtener el puerto desde .env o usar ":8000"
	// port = os.Getenv("PORT") // Cambiado a `=`
 	port = "8080"
	if port == "" {
		port = "8080" // Sin ":"
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s", port)
	r.Run(":" + port) // Aquí se asegura el formato correcto
}