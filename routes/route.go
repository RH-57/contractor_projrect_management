package routes

import (
	"projectmanagement/backend-api/controllers"
	"projectmanagement/backend-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/login", controllers.Login)

	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.GetAllUser)
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)

	return router
}
