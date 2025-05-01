package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"backend-go/auth"
	"backend-go/database"
	"backend-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login maneja la solicitud de inicio de sesión y genera un token JWT.
func Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Intento de leer los datos enviados por el cliente
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	// Depuración de las credenciales (opcional)
	log.Println("Usuario recibido:", creds.Username)
	log.Println("Contraseña recibida:", creds.Password) // NO hacerlo en producción

	// Buscar usuario en la base de datos
	var user models.User
	if err := database.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al acceder a la base de datos"})
		}
		return
	}

	log.Println("Contraseña en la BD:", user.Password)
	log.Println("Contraseña ingresada:", creds.Password)

	// Verifica si la contraseña ingresada coincide con el hash almacenado en la base de datos
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Generar el token JWT
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// Establecer el tiempo de expiración del token a 24 horas
	expiration := time.Now().Add(time.Hour * 24)

	// Enviar el token y la fecha de expiración al cliente
	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_at": expiration.Format(time.RFC3339),
	})
}

// Handler para registrar un nuevo usuario
func Register(c *gin.Context) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Obtener los datos del cuerpo de la solicitud
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validar que el nombre de usuario no esté vacío
	if userInput.Username == "" || userInput.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de usuario y la contraseña son obligatorios"})
		return
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la contraseña"})
		return
	}

	// Crear el usuario
	user := models.User{
		Username: userInput.Username,
		Password: string(hashedPassword),
		Role:     "owner", // O el valor predeterminado que prefieras
	}

	// Guardar el nuevo usuario en la base de datos
	if err := database.DB.Create(&user).Error; err != nil {
		log.Println("Error al registrar el usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el usuario"})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, gin.H{"message": "Usuario registrado con éxito"})
}