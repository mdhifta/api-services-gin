package models

import (
	"gorm.io/gorm"
)

/* Set Default Users Table */
type Users struct {
	gorm.Model

	User_id      int    `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Fname        string `json:"fname"`
	Lname        string `json:"lname"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Roles        int    `json:"roles"` // 0 = admin or 1 = users
}

// if you want to setting name table
func (users *Users) TableName() string {
	return "users"
}
