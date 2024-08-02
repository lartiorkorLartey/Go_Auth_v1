package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/InnocentEdem/Go_Auth_v1/controllers"
	_ "github.com/InnocentEdem/Go_Auth_v1/docs"
	"github.com/InnocentEdem/Go_Auth_v1/initializers"
	"github.com/InnocentEdem/Go_Auth_v1/middlewares"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvironment()
	initializers.ConnectDB()
	// initializers.SyncDatabase()
}

// @title           Gatekeeper Pro API
// @version         1.0
// @description     GateKeeper Pro registers and validates users for your frontend.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://go-auth-v1.onrender.com
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	
	r.POST("/create-client", controllers.ClientSignup)
	r.POST("/client-login", controllers.ClientLogin)
	r.POST("/generate-apn",middlewares.ClientAuthMiddleware(), controllers.GenerateAPN )
	r.POST("/invalidate-apn",middlewares.ClientAuthMiddleware(), controllers.InvalidateAPN )
	r.GET("/all-users",middlewares.ClientAuthMiddleware(), controllers.GetUsersByClient )
	r.GET("/client-apn",middlewares.ClientAuthMiddleware(), controllers.GetClientAPN )
	r.POST("/delete-user",middlewares.ClientAuthMiddleware(), controllers.DeleteUserByClient )
	r.POST("/feature-request",middlewares.ClientAuthMiddleware(), controllers.CreateFeatureRequest)
	r.GET("/config",middlewares.ClientAuthMiddleware(), controllers.GetClientAdvancedConfig)
	r.PUT("/config",middlewares.ClientAuthMiddleware(), controllers.UpdateClientAdvancedConfigHandler)
	r.GET("/client",middlewares.ClientAuthMiddleware(), controllers.GetClient)

	r.POST("/update-password",middlewares.ClientAuthMiddleware(), controllers.ClientUpdatePassword)
	
	r.POST("/user/signup", middlewares.APNAuthMiddleware(), controllers.UserSignup)
	r.POST("/user/login", middlewares.APNAuthMiddleware(), controllers.UserLogin)
	r.GET("/user/validate", middlewares.APNAuthMiddleware(), middlewares.UserAuthMiddleware(), controllers.ValidateUser)
	r.POST("/user/update-password", middlewares.APNAuthMiddleware(), middlewares.UserAuthMiddleware(), controllers.UserUpdatePassword)
	r.GET("/user/profile",middlewares.APNAuthMiddleware(), middlewares.UserAuthMiddleware(), controllers.GetUserProfile)
	r.PUT("/user/profile", middlewares.APNAuthMiddleware(), middlewares.UserAuthMiddleware(), controllers.UpdateUserProfile)

	

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run()
}