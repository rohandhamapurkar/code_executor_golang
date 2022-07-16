package models

import (
	"time"
)

type Snippets struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// disabled soft delete
	// DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name     string `json:"name" gorm:"not null;<-;unique"`
	Code     string `json:"code" gorm:"not null;<-"`
	Language string `json:"language" gorm:"not null;<-"`
	Public   bool   `json:"public" gorm:"not null;<-"`
	UserID   string `json:"userID" gorm:"not null;<-:create"`
}
