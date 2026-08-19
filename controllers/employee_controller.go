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

func GenerateEmployeeCode() (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("EMP%s", now.Format("200601"))

	var count int64
	err := database.DB.Unscoped().Model(&models.Employee{}).
		Where("code LIKE ?", prefix+"%").
		Count(&count).Error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%04d", prefix, count+1), nil
}

func GetAllEmployee(c *gin.Context) {
	var employees []models.Employee

	database.DB.Find(&employees)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Get All Employees",
		Data:    employees,
	})
}

func CreateEmployee(c *gin.Context) {
	var req = structs.EmployeeCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	employeeCode, err := GenerateEmployeeCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to generate employee code",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	employee := models.Employee{
		Code:      employeeCode,
		Name:      req.Name,
		Position:  req.Position,
		Type:      req.Type,
		Phone:     req.Phone,
		Address:   req.Address,
		DailyRate: req.DailyRate,
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Employee created successfully",
		Data: structs.EmployeeResponse{
			Id:        employee.Id,
			Code:      employee.Code,
			Name:      employee.Name,
			Position:  employee.Position,
			Type:      employee.Type,
			Phone:     employee.Phone,
			Address:   employee.Address,
			DailyRate: employee.DailyRate,
			IsActive:  employee.IsActive,
			CreatedAt: employee.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: employee.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindEmployeeById(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Employee not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Employee found",
		Data: structs.EmployeeResponse{
			Id:        employee.Id,
			Code:      employee.Code,
			Name:      employee.Name,
			Position:  employee.Position,
			Type:      employee.Type,
			Phone:     employee.Phone,
			Address:   employee.Address,
			DailyRate: employee.DailyRate,
			CreatedAt: employee.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: employee.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Employee not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req structs.EmployeeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	employee.Name = req.Name
	employee.Position = req.Position
	employee.Type = req.Type
	employee.Phone = req.Phone
	employee.Address = req.Address
	employee.DailyRate = req.DailyRate
	if req.IsActive != nil {
		employee.IsActive = *req.IsActive
	}

	if err := database.DB.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Employee updated successfully",
		Data: structs.EmployeeResponse{
			Id:        employee.Id,
			Code:      employee.Code,
			Name:      employee.Name,
			Position:  employee.Position,
			Type:      employee.Type,
			Phone:     employee.Phone,
			Address:   employee.Address,
			DailyRate: employee.DailyRate,
			CreatedAt: employee.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: employee.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Employee deleted Successfully",
	})
}
