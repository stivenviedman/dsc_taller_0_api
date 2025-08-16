package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username *string    `json:"username"`
	Password *string    `json:"password"`
	ImageP   *string    `json:"image"`
	Tasks    []Task     `json:"-"`
	Category []Category `json:"-"`
}

func MigrateUsers(db *gorm.DB) error {

	err := db.AutoMigrate(&User{})

	return err
}
