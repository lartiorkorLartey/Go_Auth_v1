package main

import (
	"fmt"

	"authapp.com/m/controllers"
	"authapp.com/m/initializers"
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
	r.Run()
}