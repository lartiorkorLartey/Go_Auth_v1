package controllers

import (
	"net/http"

	"authapp.com/m/initializers"
	"authapp.com/m/models"
	"authapp.com/m/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateAPN(c *gin.Context) {
    client, exists := c.Get("client")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    clientModel, ok := client.(models.Client)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving client information"})
        return
    }

    apn, err := utils.GenerateAPN(16)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate APN"})
        return
    }

	hash, err := bcrypt.GenerateFromPassword([]byte(apn), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate apn",
		})
		return
	}

    clientModel.APN = string(hash)

    if err := initializers.DB.Save(&clientModel).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update APN"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"apn": apn})
}

func InvalidateAPN(c *gin.Context) {
    client, exists := c.Get("client")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    clientModel, ok := client.(models.Client)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving client information"})
        return
    }

    clientModel.APN = ""

    if err := initializers.DB.Save(&clientModel).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalidate API Operation failed"})
        return
    }

    // Respond with success message
    c.JSON(http.StatusOK, gin.H{"message": "APN invalidated successfully"})
}