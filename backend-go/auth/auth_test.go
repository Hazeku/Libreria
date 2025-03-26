package auth_test

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"backend-go/auth"
)

var testJwtSecret string

func init() {
	testJwtSecret = os.Getenv("JWT_SECRET")
	if testJwtSecret == "" {
		testJwtSecret = "desarrollo_secreto_inseguro" // Debe coincidir con auth.go
	}
}

func TestGenerateJWT(t *testing.T) {
	username := "testuser"
	tokenString, err := auth.GenerateJWT(username)

	if err != nil {
		t.Fatalf("GenerateJWT failed: %v", err)
	}

	if tokenString == "" {
		t.Fatalf("GenerateJWT returned an empty token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &auth.AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(testJwtSecret), nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(*auth.AuthTokenClaims)
	if !ok {
		t.Fatalf("Failed to parse claims")
	}

	if claims.Username != username {
		t.Errorf("Expected username %s, got %s", username, claims.Username)
	}

	expirationDiff := time.Until(claims.ExpiresAt.Time)
	if expirationDiff < 59*time.Minute || expirationDiff > 61*time.Minute {
		t.Errorf("Expiration time is not within the expected range")
	}
}