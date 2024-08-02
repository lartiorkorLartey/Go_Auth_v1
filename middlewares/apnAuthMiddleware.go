package middlewares

import (
	"net/http"

	"github.com/InnocentEdem/Go_Auth_v1/initializers"
	"github.com/InnocentEdem/Go_Auth_v1/models"
	"github.com/gin-gonic/gin"
)

func APNAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        apnHeader := c.GetHeader("X-APN")

        if apnHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "APN header missing"})
            c.Abort()
            return
        }

		var client models.Client
		if err := initializers.DB.Preload("ClientAdvancedConfig").Where("apn = ?", apnHeader).First(&client).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid APN"})
			c.Abort()
			return
		}

        c.Set("client", client)
        c.Next()
    }
}
