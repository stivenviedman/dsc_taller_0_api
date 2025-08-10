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

	//pendiente campo categoria tiene que ser una relacion con la entidad category

	UserID uint `json:"user_id"`           // Clave foránea explícita
	User   User `gorm:"foreignKey:UserID"` // Relación con User
}

func MigrateTasks(db *gorm.DB) error {

	err := db.AutoMigrate(&Task{})

	return err
}
