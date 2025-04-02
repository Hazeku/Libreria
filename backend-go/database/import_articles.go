package database

import (
	"encoding/json"
	"fmt"
	"os"

	"backend-go/models"
)

// ImportArticles carga datos desde un archivo JSON y los guarda en la DB
func ImportArticles() error {
	// Leer el archivo JSON
	file, err := os.ReadFile("data/articles.json")
	if err != nil {
		return fmt.Errorf("error leyendo archivo JSON: %w", err)
	}

	// Convertir JSON a una lista de artículos
	var articles []models.Article
	err = json.Unmarshal(file, &articles)
	if err != nil {
		return fmt.Errorf("error decodificando JSON: %w", err)
	}

	// Insertar artículos en la base de datos sin duplicados
	for _, article := range articles {
		var existingArticle models.Article
		result := DB.Where("title = ?", article.Title).First(&existingArticle)

		if result.Error == nil {
			// Si ya existe, lo ignoramos
			fmt.Printf("⚠️  El artículo '%s' ya existe, omitiendo inserción.\n", article.Title)
			continue
		}

		// Si no existe, lo creamos
		if err := DB.Create(&article).Error; err != nil {
			return fmt.Errorf("error insertando artículo: %w", err)
		}
	}

	fmt.Println("✅ Artículos importados correctamente sin duplicados.")
	return nil
}
