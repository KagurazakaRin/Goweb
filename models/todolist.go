package models

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	Uid    int    `json:"uid"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	User   User   `gorm:"foreignKey:Uid"`
}
