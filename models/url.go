package models

import (
	"github.com/google/uuid"
)

type URL struct {
	Id   uuid.UUID `json:"id" gorm:"type:varchar(70);primaryKey"`
	Url  string    `json:"url" gorm:"type:varchar(40);unique;not null"`
	Site string    `json:"site" gorm:"type:varchar(40);unique;not null"`
}
