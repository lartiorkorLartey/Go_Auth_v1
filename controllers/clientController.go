package controllers

import (
	"net/http"

	"authapp.com/m/initializers"
	"authapp.com/m/models"
	"authapp.com/m/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ClientSignup(c *gin.Context) {
	var body struct {
		FirstName string `json:"firstname" binding:"required"`
		LastName  string `json:"lastname" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid properties in request body"})
		return
	}

	var client models.Client
	if err := initializers.DB.Where("email = ?", body.Email).First(&client).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with email already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating account",
		})
		return
	}

	client = models.Client{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash),
	}	
	
	result := initializers.DB.Create(&client)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating client",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": " User created successfully",
	})
}

func ClientLogin(c *gin.Context) {
    var body struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid properties in request body"})
        return
    }

    var client models.Client
    if err := initializers.DB.Where("email = ?", body.Email).First(&client).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(body.Password)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
        return
    }
	token, err := utils.GenerateJWT(body.Email, "CLIENT" )
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "login_token": token})
}