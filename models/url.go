package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey"`
	Url  string `json:"url" gorm:"type:varchar(40);unique;not null"`
	Site string `json:"site" gorm:"type:varchar(40);unique;not null"`
}
