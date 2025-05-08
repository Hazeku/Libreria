package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"backend-go/database"
	"backend-go/models"
)

// CreateArticle maneja la creación de un nuevo artículo.
func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	category := c.PostForm("category")
	priceStr := c.PostForm("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Precio inválido"})
		return
	}

	// Validar que el precio sea mayor a cero
	if price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El precio debe ser mayor a cero"})
		return
	}

	// Manejo de imagen
	file, err := c.FormFile("image")
	var imagePath string
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		fullPath := "public/Images/" + filename

		err := c.SaveUploadedFile(file, fullPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
			return
		}

		imagePath = "/Images/" + filename
	} else {
		imagePath = ""
	}

	article := models.Article{
		Title:       title,
		Description: description,
		Category:    category,
		Price:       price,
		Image:       imagePath,
	}

	// Guardar en DB
	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el artículo"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func CargarArticulosDesdeJSON() error {
	// Abrir el archivo JSON
	file, err := os.Open("data/articles.json")
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo JSON: %w", err)
	}
	defer file.Close()

	// Leer el contenido del archivo
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error al leer el archivo JSON: %w", err)
	}

	// Definir un slice de artículos
	var articulos []models.Article
	if err := json.Unmarshal(byteValue, &articulos); err != nil {
		return fmt.Errorf("error al parsear JSON: %w", err)
	}

	// Guardar los artículos en la base de datos
	for _, articulo := range articulos {
		// Verificar si el artículo ya existe
		var existingArticle models.Article
		err := database.DB.First(&existingArticle, "id = ?", articulo.ID).Error
		if err == nil {
			// El artículo ya existe, lo omitimos
			fmt.Printf("El artículo con ID %d ya existe. Se omite la inserción.\n", articulo.ID)
			continue
		}

		// Si no existe, lo insertamos
		if err := database.DB.Create(&articulo).Error; err != nil {
			return fmt.Errorf("error al guardar el artículo en la base de datos: %w", err)
		}
	}

	fmt.Println("Artículos cargados exitosamente desde el archivo JSON")
	return nil
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
	// Buscar el artículo por ID
	result := database.DB.First(&article, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Artículo con ID %d no encontrado", id),
		})
		return
	}

	database.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Artículo eliminado"})
}

// Obtener todos los artículos
func GetArticles(c *gin.Context) {
	// Lista para almacenar los artículos
	var articles []models.Article

	// Realizar la consulta sin paginación
	if err := database.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los artículos"})
		return
	}

	// Calcular el total de artículos
	var totalArticles int64
	if err := database.DB.Model(&models.Article{}).Count(&totalArticles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contar los artículos"})
		return
	}

	// Devolver todos los artículos sin paginación
	c.JSON(http.StatusOK, gin.H{
		"articles":       articles,
		"total_articles": totalArticles,
	})
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

	// Validación de los datos recibidos
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si el usuario existe
	fmt.Println("Buscando usuario con ID:", request.UserID) // Aquí agregamos el log
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
		log.Printf("Error al actualizar artículo: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo asignar el artículo"})
		return
	}

	// Verificar si se actualizó algún registro
	if result.RowsAffected == 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "El artículo no pudo ser asignado"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Artículo asignado correctamente"})
}

func UpdateArticle(c *gin.Context) {
	fmt.Println("UpdateArticle ejecutado")

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de artículo inválido"})
		return
	}

	var article models.Article
	// Buscar el artículo por ID
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Artículo con ID %d no encontrado", id),
		})
		return
	}

	// Verificar si se subió un archivo de imagen
	file, err := c.FormFile("image")

	// 👉 Agregá esta parte para ver si llega la imagen
	if err != nil {
		fmt.Println("⚠️ No se recibió ninguna imagen:", err)
	} else {
		fmt.Println("📷 Imagen recibida:", file.Filename)
	}

	// Verificar si el campo title está presente en la solicitud y actualizarlo
	title := c.PostForm("title")
	if title != "" {
		article.Title = title
	}

	// Si se subió una nueva imagen, guardarla
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		fullPath := "public/Images/" + filename
		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
			return
		}

		// Actualizar solo el campo image
		article.Image = "/Images/" + filename
	}

	// Si el campo de imagen es distinto de vacío, actualizarlo
	if article.Image != "" {
		if err := database.DB.Model(&article).Update("image", article.Image).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la imagen del artículo"})
			return
		}
	}

	// Imprimir los datos recibidos para depuración
	fmt.Println("Artículo actualizado:", article)

	// Responder con el artículo actualizado
	c.JSON(http.StatusOK, article)
}
