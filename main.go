package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"

	"github.com/InnocentEdem/goauth/controllers"
	"github.com/InnocentEdem/goauth/initializers"
	"github.com/InnocentEdem/goauth/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/InnocentEdem/goauth/docs" 

)

func init(){
	initializers.LoadEnvironment()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

// @title           Swagger Example API
// @version         1.0
// @description     GateKeeper Pro register and validates users for your frontend.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
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
	
	r.POST("/user/signup", middlewares.APNAuthMiddleware(), controllers.UserSignup)
	r.POST("/user/login", middlewares.APNAuthMiddleware(), controllers.UserLogin)
	r.GET("/user/validate", middlewares.APNAuthMiddleware(), middlewares.UserAuthMiddleware(), controllers.ValidateUser)
	r.POST("/user/update-password", middlewares.APNAuthMiddleware(), middlewares.UserAuthMiddleware(), controllers.UserUpdatePassword)
	

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run()
}