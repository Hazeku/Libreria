package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Definición de la estructura para el token JWT
type AuthTokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtSecret = os.Getenv("JWT_SECRET")

func init() {
	if jwtSecret == "" {
		fmt.Println("WARNING: JWT_SECRET environment variable not set. Using a default for development.")
		jwtSecret = "desarrollo_secreto_inseguro" // ¡No usar esto en producción!
	}
}

// Función para generar un token JWT
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &AuthTokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}

// Función para validar un token JWT
func ValidateJWT(tokenString string) (string, error) {
	claims := &AuthTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", jwt.ErrInvalidKey
	}
	return claims.Username, nil
}

// Función para verificar si una contraseña coincide con un hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Función para hashear una contraseña (si la necesitas en este paquete)
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
