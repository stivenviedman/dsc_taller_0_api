package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Description      *string   `json:"description"`
	CreationDate     time.Time `gorm:"type:date" json:"creationDate"`     //formato RFC3339
	FinalizationDate time.Time `gorm:"type:date" json:"finalizationDate"` //formato RFC3339
	State            *string   `json:"state"`

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
