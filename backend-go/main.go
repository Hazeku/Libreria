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
)

func GetClientIP(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		return strings.Split(forwardedFor, ",")[0]
	}
	return r.RemoteAddr
}

func myHandler(c *gin.Context) {
	ginClientIP := c.ClientIP()

	c.JSON(http.StatusOK, gin.H{
		"ginClientIP": ginClientIP,
		"message":     "Hello from handler",
	})
}

func loadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in the .env file")
	}

	fmt.Println("JWT_SECRET: ", jwtSecret)
	config.LoadConfig()
}

func main() {
	loadConfiguration()

	r := gin.Default()

	// 🔒 Servir archivos estáticos para imágenes
	r.Static("/Images", "./public/Images")

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.1.40:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.SetTrustedProxies([]string{
		"192.168.1.45",
		"192.168.0.1",
	})

	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}

	if err := handlers.CargarArticulosDesdeJSON(); err != nil {
		log.Fatalf("Error cargando artículos desde JSON: %v", err)
	}

	if os.Getenv("ENV") != "production" {
		if err := database.ImportArticles(); err != nil {
			log.Printf("⚠️  No se pudieron importar los artículos: %v", err)
		}
	}

	// 🌐 Rutas públicas
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	r.GET("/articles", handlers.GetArticles) // ✅ Ruta pública para el PublicHome
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Bienvenido a la API"})
	})
	r.GET("/hello", myHandler)

	// 🔐 Rutas protegidas
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("/articles", handlers.CreateArticle)
		authGroup.DELETE("/articles/:id", handlers.DeleteArticle)
		authGroup.PUT("/articles/:id/assign", handlers.AssignArticleToUser)
		authGroup.PUT("/articles/:id", handlers.UpdateArticle)
	}

	// Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s", port)
	r.Run(":" + port)
}
