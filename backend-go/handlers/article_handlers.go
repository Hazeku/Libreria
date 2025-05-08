package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"backend-go/database"
	"backend-go/models"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	category := c.PostForm("category")
	priceStr := c.PostForm("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Precio inválido"})
		return
	}

	// Manejo de imagen
	file, err := c.FormFile("image")
	var imagePath string
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		fullPath := "public/Images/" + filename
		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
			return
		}
		imagePath = "/Images/" + filename
	}

	article := models.Article{
		Title:       title,
		Description: description,
		Category:    category,
		Price:       price,
		Image:       imagePath,
	}

	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el artículo"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artículo no encontrado"})
		return
	}

	// Actualización de campos si están presentes
	if title := c.PostForm("title"); title != "" {
		article.Title = title
	}
	if description := c.PostForm("description"); description != "" {
		article.Description = description
	}
	if category := c.PostForm("category"); category != "" {
		article.Category = category
	}
	if priceStr := c.PostForm("price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			article.Price = price
		}
	}

	// Imagen
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		fullPath := "public/Images/" + filename
		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
			return
		}
		article.Image = "/Images/" + filename
	}

	if err := database.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el artículo"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func GetArticles(c *gin.Context) {
	var articles []models.Article
	if err := database.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los artículos"})
		return
	}

	var total int64
	database.DB.Model(&models.Article{}).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"articles":       articles,
		"total_articles": total,
	})
}

func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artículo no encontrado"})
		return
	}

	database.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Artículo eliminado"})
}

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

	var user models.User
	if err := database.DB.First(&user, request.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	var article models.Article
	if err := database.DB.First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artículo no encontrado"})
		return
	}

	if article.AuthorID != nil && *article.AuthorID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Ya asignado a otro usuario"})
		return
	}

	if err := database.DB.Model(&article).Update("author_id", request.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo asignar el artículo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artículo asignado correctamente"})
}

func CargarArticulosDesdeJSON() error {
	file, err := os.Open("data/articles.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error al leer el archivo JSON: %w", err)
	}

	var articulos []models.Article
	if err := json.Unmarshal(byteValue, &articulos); err != nil {
		return fmt.Errorf("error al parsear JSON: %w", err)
	}

	for _, articulo := range articulos {
		var existing models.Article
		if err := database.DB.First(&existing, "id = ?", articulo.ID).Error; err == nil {
			continue
		}
		if err := database.DB.Create(&articulo).Error; err != nil {
			return fmt.Errorf("error al guardar artículo: %w", err)
		}
	}
	return nil
}
