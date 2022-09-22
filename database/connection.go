package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goweb/models"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := "root:windows@(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	err = DB.AutoMigrate(&models.TodoList{}, &models.User{})
	if err != nil {
		panic("migrate failed")
	}

}
