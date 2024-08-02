package utils

import (
	"github.com/InnocentEdem/Go_Auth_v1/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)


func SetDefaultClientAdvancedConfig(clientId uuid.UUID) models.ClientAdvancedConfig {
	return models.ClientAdvancedConfig{
		ClientID: clientId,
		CorsAllowedOrigins:    pq.StringArray{""},
		JWTExpiryTime:         3600,
		RefreshTokenEnabled:   true,
		RefreshTokenExpiryTime: 7200,
		AllowJWTCustomClaims:  false,  
		UseAdditionalProperties: false,    
	}
}