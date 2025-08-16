package models

import (
	"gorm.io/gorm"
)

type Category struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name" gorm:"not null;unique"`
	Description string `json:"description" gorm:"not null"`

	// Relaciones
	Tasks []Task `json:"tasks,omitempty"`
	// Relaciones
	UserID uint `json:"user_id"` // Clave foránea explícita
	User   User `gorm:"foreignKey:UserID"`
}

func MigrateCategories(db *gorm.DB) error {
	err := db.AutoMigrate(&Category{})
	return err
}
