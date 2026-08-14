package routes

import (
	"projectmanagement/backend-api/controllers"
	"projectmanagement/backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/api/login", controllers.Login)

	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.GetAllUser)
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	return router
}
