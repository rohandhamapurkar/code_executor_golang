package models

import (
	"gorm.io/gorm"
)

type Snippets struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;<-;unique"`
	Code     string `json:"code" gorm:"not null;<-"`
	Language string `json:"language" gorm:"not null;<-"`
	Public   bool   `json:"public" gorm:"not null;<-"`
	UserID   string `json:"userID" gorm:"not null;<-:create"`
}
