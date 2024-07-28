package main

import (
	"fmt"

	"authapp.com/m/controllers"
	"authapp.com/m/initializers"
	"authapp.com/m/middlewares"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvironment()
	initializers.ConnectDB()
	// initializers.SyncDatabase()
}

func main() {
	fmt.Println("working 1")
	r := gin.Default()
	r.POST("/create-client", controllers.ClientSignup)
	r.POST("/client-login", controllers.ClientLogin)
	r.POST("/generate-apn",middlewares.ClientAuthMiddleware(), controllers.GenerateAPN )
	r.POST("/invalidate-apn",middlewares.ClientAuthMiddleware(), controllers.InvalidateAPN )
	r.Run()
}