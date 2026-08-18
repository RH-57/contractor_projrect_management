package controllers

import (
	"fmt"
	"net/http"
	"projectmanagement/backend-api/database"
	"projectmanagement/backend-api/helpers"
	"projectmanagement/backend-api/models"
	"projectmanagement/backend-api/structs"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllProject(c *gin.Context) {
	var projects []models.Project

	database.DB.Preload("Customer").Find(&projects)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Get All Projects",
		Data:    projects,
	})
}

func GenerateProjectCode() (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("PRJ-%s-", now.Format("200601"))

	var count int64

	err := database.DB.Model(&models.Project{}).
		Where("code LIKE ?", prefix+"%").
		Count(&count).Error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%04d", prefix, count+1), nil
}

func CreateProject(c *gin.Context) {
	var req = structs.ProjectCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var customer models.Customer
	if err := database.DB.First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Customer not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  map[string]string{"start_date": "Invalid date format, use YYYY-MM-DD"},
		})

	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  map[string]string{"end_date": "Invalid date format, use YYYY-MM-DD"},
		})
		return
	}

	projectCode, err := GenerateProjectCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to generate project code",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	status := req.Status
	if status == "" {
		status = "PLANNED"
	}

	project := models.Project{
		Code:          projectCode,
		Name:          req.Name,
		CustomerID:    req.CustomerID,
		ContractValue: req.ContractValue,
		EstimatedCost: req.EstimatedCost,
		Status:        status,
		StartDate:     startDate,
		EndDate:       endDate,
	}

	// 7. Simpan ke Database
	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// 8. Response Sukses
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Project created successfully",
		Data: structs.ProjectResponse{
			Id:            project.Id,
			Code:          project.Code,
			Name:          project.Name,
			CustomerID:    project.Customer.Id,
			ContractValue: project.ContractValue,
			EstimatedCost: project.EstimatedCost,
			Status:        project.Status,
			StartDate:     project.StartDate.Format("2006-01-02"),
			EndDate:       project.EndDate.Format("2006-01-02"),
			CreatedAt:     project.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     project.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
