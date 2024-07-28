package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    FirstName string `gorm:"type:varchar(100);not null"`
    LastName  string `gorm:"type:varchar(100);not null"`
    Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
    Password  string `gorm:"type:varchar(255);not null"`
    APN       string `gorm:"type:varchar(100);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
    Users     []User
}

func (c *Client) BeforeCreate(tx *gorm.DB) (err error) {
    c.ID = uuid.New()
    return
}
