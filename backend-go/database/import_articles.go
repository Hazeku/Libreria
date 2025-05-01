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
		result := DB.FirstOrCreate(&existingArticle, models.Article{Title: article.Title})

		if result.Error != nil {
			return fmt.Errorf("error insertando artículo: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			// Si no se insertó nada, significa que ya existía
			fmt.Printf("⚠️  El artículo '%s' ya existe, omitiendo inserción.\n", article.Title)
		} else {
			// Si el artículo fue insertado, lo mostramos
			fmt.Printf("✅ Artículo '%s' insertado correctamente.\n", article.Title)
		}
	}

	fmt.Println("✅ Artículos importados correctamente sin duplicados.")
	return nil
}
