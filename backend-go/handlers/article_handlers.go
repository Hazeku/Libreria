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

// Obtener todos los artículos
func GetArticles(c *gin.Context) {
	var articles []models.Article

	// Consulta a la base de datos
	if err := database.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los artículos"})
		return
	}

	// Responder con la lista de artículos
	c.JSON(http.StatusOK, articles)
}

// AssignArticleToUser asigna un artículo a un usuario cuando se compra
func AssignArticleToUser(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	var request struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si el usuario existe
	var user models.User
	if err := database.DB.First(&user, request.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Verificar si el artículo existe
	var article models.Article
	if err := database.DB.First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artículo no encontrado"})
		return
	}

	// Verificar si el artículo ya está asignado
	if article.AuthorID != nil && *article.AuthorID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "El artículo ya está asignado a otro usuario"})
		return
	}

	// Asignar el artículo al usuario
	result := database.DB.Model(&article).Update("author_id", request.UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo asignar el artículo"})
		return
	}

	if result.Error != nil {
		log.Printf("Error al actualizar artículo: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo asignar el artículo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artículo asignado correctamente"})
}
