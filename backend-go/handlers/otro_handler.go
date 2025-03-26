package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLibraryInfo maneja la solicitud para obtener información sobre la librería.
func GetLibraryInfo(c *gin.Context) {
	libraryInfo := map[string]string{
		"name":    "Mi Librería",
		"address": "Calle Falsa 123",
		"phone":   "555-1234",
	}
	c.JSON(http.StatusOK, libraryInfo)
}