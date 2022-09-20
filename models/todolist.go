package models

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
