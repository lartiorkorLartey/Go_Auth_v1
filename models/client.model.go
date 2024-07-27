package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	APN  	  string `gorm:"type:varchar(100);uniqueIndex"`
	Users     []User
}