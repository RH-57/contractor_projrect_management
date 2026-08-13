package routes

import (
	"projectmanagement/backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/login", controllers.Login)

	return router
}
