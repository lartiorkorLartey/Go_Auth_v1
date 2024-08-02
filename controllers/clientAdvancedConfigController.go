package controllers

import (
	"net/http"

	"github.com/InnocentEdem/Go_Auth_v1/initializers"
	"github.com/InnocentEdem/Go_Auth_v1/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)



func UpdateClientAdvancedConfigHandler(c *gin.Context) {
	client, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing client id"})
		return
	}

	clientModel, ok := client.(models.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving client information"})
		return
	}

	var req UpdateClientAdvancedConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var config models.ClientAdvancedConfig
	if err := initializers.DB.Where("client_id = ?", clientModel.ID).First(&config).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client advanced config not found"})
		return
	}

	if req.CorsAllowedOrigins != nil {
		config.CorsAllowedOrigins = *req.CorsAllowedOrigins
	}

	if req.JWTExpiryTime != nil {
		config.JWTExpiryTime = *req.JWTExpiryTime
	}
	if req.RefreshTokenEnabled != nil {
		config.RefreshTokenEnabled = *req.RefreshTokenEnabled
	}
	if req.RefreshTokenExpiryTime != nil {
		config.RefreshTokenExpiryTime = *req.RefreshTokenExpiryTime
	}
	if req.AllowJWTCustomClaims != nil {
		config.AllowJWTCustomClaims = *req.AllowJWTCustomClaims
	}
	if req.UseAdditionalProperties != nil {
		config.UseAdditionalProperties = *req.UseAdditionalProperties
	}

	if err := initializers.DB.Save(&config).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client advanced config"})
		return
	}

	response := ClientAdvancedConfigResponse{
		ID:                    config.ID,
		ClientID:              config.ClientID,
		CorsAllowedOrigins:    config.CorsAllowedOrigins,
		JWTExpiryTime:         config.JWTExpiryTime,
		RefreshTokenEnabled:   config.RefreshTokenEnabled,
		RefreshTokenExpiryTime: config.RefreshTokenExpiryTime,
		AllowJWTCustomClaims:  config.AllowJWTCustomClaims,
		UseAdditionalProperties: config.UseAdditionalProperties,
	}

	c.JSON(http.StatusOK, response)
}

func GetClientAdvancedConfig(c *gin.Context) {

	client, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing client id"})
		return
	}

	clientModel, ok := client.(models.Client)
    if !ok {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error retrieving client information"})
        return
    }


	var clientConfig models.ClientAdvancedConfig
	if err := initializers.DB.First(&clientConfig, "client_id = ?", clientModel.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client configuration not found"})
		return
	}

	response := ClientAdvancedConfigResponse{
		ID:                    clientConfig.ID,
		ClientID:              clientConfig.ClientID,
		CorsAllowedOrigins:    clientConfig.CorsAllowedOrigins,
		JWTExpiryTime:         clientConfig.JWTExpiryTime,
		RefreshTokenEnabled:   clientConfig.RefreshTokenEnabled,
		RefreshTokenExpiryTime: clientConfig.RefreshTokenExpiryTime,
		AllowJWTCustomClaims:  clientConfig.AllowJWTCustomClaims,
		UseAdditionalProperties: clientConfig.UseAdditionalProperties,
	}

	c.JSON(http.StatusOK, response)
}

type ClientAdvancedConfigResponse struct {
    ID                    uuid.UUID `json:"id"`
	ClientID              uuid.UUID `json:"client_id"`
	CorsAllowedOrigins    []string `json:"cors_allowed_origins"`
	JWTExpiryTime         int      `json:"jwt_expiry_time"`
	RefreshTokenEnabled   bool     `json:"refresh_token_enabled"`
	RefreshTokenExpiryTime int     `json:"refresh_token_expiry_time"`
	AllowJWTCustomClaims  bool     `json:"allow_jwt_custom_claims"`
	UseAdditionalProperties  bool  `json:"use_additional_properties"`
}
type UpdateClientAdvancedConfigRequest struct {
	CorsAllowedOrigins      *[]string `json:"cors_allowed_origins,omitempty"`
	JWTExpiryTime           *int      `json:"jwt_expiry_time,omitempty"`
	RefreshTokenEnabled     *bool     `json:"refresh_token_enabled,omitempty"`
	RefreshTokenExpiryTime  *int      `json:"refresh_token_expiry_time,omitempty"`
	AllowJWTCustomClaims    *bool     `json:"allow_jwt_custom_claims,omitempty"`
	UseAdditionalProperties *bool     `json:"use_additional_properties,omitempty"`
}