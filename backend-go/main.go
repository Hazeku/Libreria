package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	InitDB()

	// Rutas públicas
	r.POST("/login", login)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Bienvenido a la API"})
	})

	// Rutas protegidas (solo para el propietario)
	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	{
		auth.POST("/articles", createArticle)
		auth.DELETE("/articles/:id", deleteArticle)
	}

	r.Run(":8000")
}

// Definir al usuario propietario (esto debería moverse a un archivo de configuración)
var ownerUser = User{
	Username: "owner",
	Password: "$2a$14$2b2yAq3TpUMzJ/aWzQ/uX.WJ8gUwhcOj5dE4zOxoZy.Bc/lXm9Jde", // Contraseña hasheada (ejemplo)
}

// Handler para iniciar sesión y obtener un token
func login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Validar usuario
	if creds.Username != ownerUser.Username || !CheckPasswordHash(creds.Password, ownerUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	token, _ := GenerateJWT(ownerUser.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Middleware de autenticación
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			c.Abort()
			return
		}

		username, err := ValidateJWT(tokenString)
		if err != nil || username != ownerUser.Username {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Crear un nuevo artículo
func createArticle(c *gin.Context) {
	var article Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	DB.Create(&article)
	c.JSON(http.StatusOK, article)
}

// Eliminar un artículo por ID
func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article Article
	result := DB.First(&article, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artículo no encontrado"})
		return
	}

	DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Artículo eliminado"})
}
