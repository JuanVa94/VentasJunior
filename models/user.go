package models

import (
	"github.com/jinzhu/gorm"
)

// User es la estrucutra del usuario
type User struct {
	gorm.Model
	Username        string  `json:"username" gorm:"not null;unique"`
	Email           string  `json:"email" gorm:"not null;unique"`
	Fullname        string  `json:"fullname" gorm:"not nullt"`
	Password        string  `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string  `json:"confirmPassword,omitempty" gorm:"-"`
	Picture         string  `json:"picture"`
	Orders          []Order `json:"orders,omitempty"`
}
