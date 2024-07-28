package middlewares

import (
	"net/http"
	"strings"

	"authapp.com/m/initializers"
	"authapp.com/m/models"
	"authapp.com/m/utils"
	"github.com/gin-gonic/gin"
)

func ClientAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
            c.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        
        claims, err := utils.ParseJWTWithClaims(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        var client models.Client
        if err := initializers.DB.Where("email = ?", claims.Email).First(&client).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid client"})
            c.Abort()
            return
        }

        c.Set("client", client)
        c.Next()
    }
}