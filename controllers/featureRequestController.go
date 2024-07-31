package controllers

import (
	"net/http"

	"github.com/InnocentEdem/goauth/initializers"
	"github.com/InnocentEdem/goauth/models"
	"github.com/gin-gonic/gin"
)



func CreateFeatureRequest(c *gin.Context) {
    var input FeatureRequestInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    client, exists := c.Get("client")
    if !exists {
        c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Client not authenticated"})
        return
    }

    clientModel, ok := client.(models.Client)
    if !ok {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error retrieving client information"})
        return
    }

    featureRequest := models.FeatureRequest{
        Feature: input.Feature,
        Title:   input.Title,
		Email: clientModel.Email,
    }

    if err := initializers.DB.Create(&featureRequest).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feature request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Feature request submitted successfully"})
}

type FeatureRequestInput struct {
    Feature string `json:"feature" binding:"required"`
    Title   string `json:"title" binding:"required"`
}