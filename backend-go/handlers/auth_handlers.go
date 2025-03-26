package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend-go/auth"
	"backend-go/models"
)

// Login maneja la solicitud de inicio de sesión y genera un token JWT.
func Login(owner models.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "details": err.Error()})
			return
		}

		// Validar usuario
		if creds.Username != owner.Username || !auth.CheckPasswordHash(creds.Password, owner.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
			return
		}

		token, _ := auth.GenerateJWT(owner.Username)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}