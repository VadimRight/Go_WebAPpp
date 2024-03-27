package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"type:varchar(40);unique;not null"`
	Email    string `json:"email" gorm:"type:varcahr(40);uniqueIndex;not null"`
	Password string `gorm:"syze:255;" json:"password,omitempty"`
}
