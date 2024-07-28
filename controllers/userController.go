package controllers

import (
	"net/http"

	"authapp.com/m/initializers"
	"authapp.com/m/models"
	"authapp.com/m/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(c *gin.Context) {
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

    var existingUser models.User
    if err := initializers.DB.Where("client_id = ? AND email = ?", clientModel.ID, body.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists for this client"})
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating account"})
        return
    }

    user := models.User{
        FirstName: body.FirstName,
        LastName:  body.LastName,
        Email:     body.Email,
        Password:  string(hash),
        ClientID:  clientModel.ID,
    }

    result := initializers.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating client"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}


func UserLogin(c *gin.Context) {
    var body struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid properties in request body"})
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

    var user models.User
    if err := initializers.DB.Where("client_id = ? AND email = ?", clientModel.ID, body.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := utils.GenerateUserJWT(body.Email, "CLIENT", user.ClientID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "login_token": token})
}

func ValidateUser(c *gin.Context) {
	user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing client id"})
        return
    }

	_, ok := user.(models.User)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
	}

	c.JSON(http.StatusOK, gin.H{"message":"okay"})
}
