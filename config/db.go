package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	//dsn := "root:mysql123@tcp(127.0.0.1:3307)/contactlist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE")), &gorm.Config{})

	return db, err
}
