package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeatureRequest struct {
	gorm.Model
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Feature   string   `gorm:"type:varchar(100);not null"`
    Title     string   `gorm:"type:varchar(100);not null"`
    Email     string   `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    return
}
