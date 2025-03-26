package handlers
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"backend-go/database" /* ruta */
	"backend-go/models"
)

func TestCreateArticleIntegration(t *testing.T) {
	// Configurar un entorno de prueba (base de datos en memoria, etc.)
	os.Setenv("GIN_MODE", "test")
	gin.SetMode(gin.TestMode)

	// Inicializar la base de datos de prueba
	testDB, err := database.InitTestDB() // Función para inicializar una base de datos de prueba
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	database.DB = testDB
	defer database.CloseTestDB(testDB) // Cerrar la base de datos de prueba al finalizar

	// Crear un router de Gin para las pruebas
	r := gin.Default()
	r.POST("/articles", CreateArticle) // Usar CreateArticle directamente

	// Crear los datos del artículo a enviar
	articleData := models.Article{
		Title:   "Test Article",
		Content: "This is a test article.",
	}
	body, _ := json.Marshal(articleData)

	// Crear una solicitud HTTP
	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Crear un recorder para capturar la respuesta
	w := httptest.NewRecorder()

	// Realizar la solicitud al handler
	r.ServeHTTP(w, req)

	// Verificar la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	var createdArticle models.Article
	err = json.Unmarshal(w.Body.Bytes(), &createdArticle)
	assert.NoError(t, err)
	assert.Equal(t, articleData.Title, createdArticle.Title)
	assert.Equal(t, articleData.Content, createdArticle.Content)

	// Verificar que el artículo se guardó en la base de datos
	var retrievedArticle models.Article
	result := database.DB.First(&retrievedArticle, createdArticle.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, articleData.Title, retrievedArticle.Title)
}