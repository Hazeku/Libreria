package middleware

import (
	"backend-go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware verifica si el usuario está autenticado utilizando JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token JWT de las cabeceras Authorization (por ejemplo, "Bearer <token>")
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		// El token usualmente se pasa con el prefijo "Bearer "
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:] // Extraer el token real
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		// Verificar la validez del token y obtener el username desde el token
		username, err := auth.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Almacenar el username en el contexto de la solicitud para usarlo en las rutas protegidas
		c.Set("username", username)

		// Continuar con la solicitud
		c.Next()
	}
}
