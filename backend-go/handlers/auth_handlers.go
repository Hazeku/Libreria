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

	// Verificar si la contraseña coincide con el hash almacenado
	// if !auth.CheckPasswordHash(creds.Password, user.Password) {
	if creds.Password != user.Password {
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
