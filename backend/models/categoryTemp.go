package models

import (
	"gorm.io/gorm"
)

type CategoryTemp struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`

	Tasks []Task `json:"-"`
}

func MigrateCategories(db *gorm.DB) error {

	err := db.AutoMigrate(&CategoryTemp{})

	return err
}
