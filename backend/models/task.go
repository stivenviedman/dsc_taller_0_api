package models

import (
	"gorm.io/gorm"
)

type Task struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Description *string `json:"description"`
	Date        *string `json:"date"`

	// Relaciones
	UserID     uint     `json:"user_id"`                     // Clave foránea explícita
	User       User     `gorm:"foreignKey:UserID"`           // Relación con User
	CategoryID uint     `json:"category_id" gorm:"not null"` // Clave foránea obligatoria
	Category   Category `gorm:"foreignKey:CategoryID"`       // Relación con Category
}

func MigrateTasks(db *gorm.DB) error {
	err := db.AutoMigrate(&Task{})
	return err
}
