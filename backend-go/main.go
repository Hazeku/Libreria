package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend-go/config"
	"backend-go/database"
	"backend-go/handlers"
	"backend-go/middleware"
	// "golang.org/x/crypto/bcrypt"
)

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
	// llamar a la funcion de variables de entorno
	loadConfiguration()

	// Inicializar Gin
	r := gin.Default()

	// Habilitar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // orígenes permitidos
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Permitir enviar cookies
	}))

	// Configurar la confianza de proxies
	r.SetTrustedProxies([]string{
		"192.168.1.45", // IP local de tu servidor
		"192.168.0.1",  // IP de tu proxy
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
	r.POST("/register", handlers.Register)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Bienvenido a la API"})
	})
	r.GET("/hello", myHandler)

	// Rutas protegidas
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware()) // Middleware de autenticación JWT
	{
		authGroup.GET("/articles", handlers.GetArticles)
		authGroup.POST("/articles", handlers.CreateArticle)
		authGroup.DELETE("/articles/:id", handlers.DeleteArticle)
		authGroup.PUT("/articles/:id/assign", handlers.AssignArticleToUser)
	}

	// Configuración del puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s", port)

	// 🔐 Solo para generar el hash de "123456" UNA VEZ
	/* password := "123456"
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error al hashear la contraseña: %v", err)
	}
	fmt.Println("🔑 Hash generado para '123456':", string(hashed)) */

	// Ejecutar el servidor
	r.Static("/Images", "./public/Images")

	r.Run(":" + port)
}
