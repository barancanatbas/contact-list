package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Phone       string `gorm:"char(11)"`
	Description string `gorm:"text"`
}
