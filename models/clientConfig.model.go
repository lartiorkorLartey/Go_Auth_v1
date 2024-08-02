package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ClientAdvancedConfig struct {
	ID                    uuid.UUID           `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ClientID              uuid.UUID           `gorm:"type:uuid;unique"`
	CorsAllowedOrigins    pq.StringArray            `json:"cors_allowed_origins" gorm:"type:text[]"`
	JWTExpiryTime         int                 `json:"jwt_expiry_time"`
	RefreshTokenEnabled   bool                `json:"refresh_token_enabled"`
	RefreshTokenExpiryTime int                `json:"refresh_token_expiry_time"`
	AllowJWTCustomClaims  bool                `json:"allow_jwt_custom_claims"`
	UseAdditionalProperties  bool             `json:"use_additional_properties"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt      `gorm:"index"`
}

func (Client) TableName() string {
	return "clients"
}

func (ClientAdvancedConfig) TableName() string {
	return "client_advanced_configs"
}