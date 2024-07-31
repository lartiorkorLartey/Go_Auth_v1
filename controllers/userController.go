package controllers

import (
	"net/http"

	"authapp.com/m/initializers"
	"authapp.com/m/models"
	"authapp.com/m/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserSignup godoc
// @Summary User signup
// @Description Creates a new user account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body UserSignupRequest true "Signup details"
// @Success 200 {object} UserSignupResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/signup [post]
func UserSignup(c *gin.Context) {
    var body UserSignupRequest
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

// UserLogin godoc
// @Summary User login
// @Description Logs in a user and returns a JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body UserLoginRequest true "Login details"
// @Success 200 {object} UserLoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/login [post]
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

    token, err := utils.GenerateUserJWT(user, "User")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "login_token": token})
}

// ValidateUser godoc
// @Summary Validate user
// @Description Validates the user and checks if they are authenticated
// @Tags validate user for protected routes
// @Accept  json
// @Produce  json
// @Success 200 {object} ValidateUserSuccessResponse
// @Failure 401 {object} ValidateUserErrorResponse
// @Router /user/validate [get]
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


// UserUpdatePassword godoc
// @Summary Update user password
// @Description Updates a user's password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body  UserPasswordUpdateRequest true "Password update details"
// @Success 200 {object} UserPasswordUpdateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/update-password [post]
func UserUpdatePassword(c *gin.Context) {
    var request UserPasswordUpdateRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Missing or invalid properties in request body"})
        return
    }

    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "User not authenticated"})
        return
    }

    userModel, ok := user.(models.User)
    if !ok {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error retrieving user information"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(request.OldPassword)); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Incorrect old password"})
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error updating password"})
        return
    }

    userModel.Password = string(hash)
    if err := initializers.DB.Save(&userModel).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error updating password"})
        return
    }

    c.JSON(http.StatusOK, UserPasswordUpdateResponse{Message: "Password updated successfully"})
}

type UserSignupRequest struct {
    FirstName string `json:"firstname" binding:"required"`
    LastName  string `json:"lastname" binding:"required"`
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required"`
}
type UserLoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
    Message    string `json:"message"`
    LoginToken string `json:"login_token"`
}
type UserPasswordUpdateRequest struct {
    OldPassword string `json:"old_password" binding:"required"`
    NewPassword string `json:"new_password" binding:"required"`
}

// UserPasswordUpdateResponse represents the response for a successful password update
type UserPasswordUpdateResponse struct {
    Message string `json:"message"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

type UserSignupResponse struct {
    Message string `json:"message"`
}

type ValidateUserSuccessResponse struct {
    Message string `json:"message"`
}

type ValidateUserErrorResponse struct {
    Error string `json:"error"`
}