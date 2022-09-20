package models

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"` // `json:"-"`
}
