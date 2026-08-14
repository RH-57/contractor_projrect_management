package controllers

import (
	"net/http"
	"projectmanagement/backend-api/database"
	"projectmanagement/backend-api/models"
	"projectmanagement/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func GetAllProject(c *gin.Context) {
	var projects []models.Project

	database.DB.Find(&projects)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Get All Projects",
		Data:    projects,
	})
}
