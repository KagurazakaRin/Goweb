package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"` // `json:"-"`
	RoleID   int    `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}

const passwordCost = 14

// SetPassword error 处理 ?
func (user *User) SetPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	user.Password = string(hashPassword)
}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}
