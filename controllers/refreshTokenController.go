package controllers

import (
	"net/http"

	"github.com/InnocentEdem/Go_Auth_v1/initializers"
	"github.com/InnocentEdem/Go_Auth_v1/models"
	"github.com/InnocentEdem/Go_Auth_v1/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Refresh JWT
// @Description Refreshes the JWT and optionally the refresh token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   refresh_token  body  string  true  "Refresh Token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/refresh-token [post]
func RefreshToken(c *gin.Context) {
    var body struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid properties in request body"})
        return
    }

	user, err := utils.ParseUserRefreshJWT(body.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	
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

	if !clientModel.ClientAdvancedConfig.RefreshTokenEnabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token not activated for client"})
        return
	}


	var userData models.User
    if err := initializers.DB.Where("client_id = ? AND email = ?", clientModel.ID, user.Email).First(&userData).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
        return
    }

    newToken, err := utils.GenerateUserJWT(userData, "User")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new token"})
        return
    }
	newRefreshToken, err := utils.GenerateRefreshJWT(userData, "User")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new refresh token"})
		return
	}

	response := gin.H{
        "message": "Successful",
        "login_token": newToken,
        "refresh_token": newRefreshToken,

    }

    c.JSON(http.StatusOK, response)
}
