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

	CategoryTempID uint         `json:"category_id"`
	Category       CategoryTemp `gorm:"foreignKey:CategoryTempID"`
	UserID         uint         `json:"user_id"`
	User           User         `gorm:"foreignKey:UserID"`
}

func MigrateTasks(db *gorm.DB) error {

	err := db.AutoMigrate(&Task{})

	return err
}
