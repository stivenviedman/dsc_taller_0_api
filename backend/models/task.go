package models

import (
	"gorm.io/gorm"
)

type Task struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Description *string `json:"description"`
	Date        *string `json:"date"`
}

func MigrateTasks(db *gorm.DB) error {

	err := db.AutoMigrate(&Task{})

	return err
}
