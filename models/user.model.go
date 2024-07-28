package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    FirstName string   `gorm:"type:varchar(100);not null"`
    LastName  string   `gorm:"type:varchar(100);not null"`
    Email     string   `gorm:"type:varchar(100);not null"`
    Password  string   `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
    ClientID  uuid.UUID `gorm:"type:uuid"`
    Client    Client
}

func (User) TableName() string {
    return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    return
}

func (User) AfterMigrate(tx *gorm.DB) error {
    return tx.Exec("ALTER TABLE users ADD CONSTRAINT unique_user_client_email UNIQUE (client_id, email)").Error
}
