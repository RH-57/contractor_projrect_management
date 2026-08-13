package controllers

import (
	"net/http"
	"projectmanagement/backend-api/database"
	"projectmanagement/backend-api/helpers"
	"projectmanagement/backend-api/models"
	"projectmanagement/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Get All Users",
		Data:    users,
	})
}

func CreateUser(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User Created Successfully",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			IsActive:  true,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
