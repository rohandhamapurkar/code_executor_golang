package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"<-"`
	Email    string `gorm:"->"`
	Password string `gorm:"->:false;<-:update"`
}
