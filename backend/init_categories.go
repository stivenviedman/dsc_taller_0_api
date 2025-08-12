package main

import (
	"back-end-todolist/models"
	"log"

	"gorm.io/gorm"
)

// InitDefaultCategories crea categorías por defecto si no existen
func InitDefaultCategories(db *gorm.DB) error {
	defaultCategories := []models.Category{
		{Name: "Personal"},
		{Name: "Trabajo"},
		{Name: "Estudio"},
		{Name: "Hogar"},
		{Name: "Salud"},
	}

	for _, category := range defaultCategories {
		// Verificar si la categoría ya existe
		var existingCategory models.Category
		if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			// La categoría no existe, crearla
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Error creando categoría %s: %v", category.Name, err)
				return err
			}
			log.Printf("Categoría creada: %s", category.Name)
		}
	}

	log.Println("Categorías por defecto inicializadas correctamente")
	return nil
}
