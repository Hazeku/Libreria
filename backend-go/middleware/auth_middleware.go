package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend-go/auth"
	"backend-go/models"
)

// Middleware de autenticación
func AuthMiddleware(owner models.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			c.Abort()
			return
		}

		username, err := auth.ValidateJWT(tokenString)
		if err != nil || username != owner.Username {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado"})
			c.Abort()
			return
		}

		c.Next()
	}
}