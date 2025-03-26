package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend-go/database"
	"backend-go/models"
)

// CreateArticle maneja la creación de un nuevo artículo.
func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		log.Printf("Error binding JSON for createArticle: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos del artículo"})
		return
	}

	database.DB.Create(&article)
	c.JSON(http.StatusOK, article)
}

// DeleteArticle maneja la eliminación de un artículo por ID.
func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	var article models.Article
	result := database.DB.First(&article, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artículo no encontrado"})
		return
	}

	database.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Artículo eliminado"})
}