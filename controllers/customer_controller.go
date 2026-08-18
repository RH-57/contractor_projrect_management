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

func GetAllCustomer(c *gin.Context) {
	var customers []models.Customer

	database.DB.Find(&customers)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Get All Customers",
		Data:    customers,
	})
}

func GenerateCustomerCode() (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("CS-%s-", now.Format("200601"))

	var count int64

	err := database.DB.Model(&models.Customer{}).
		Where("code LIKE ?", prefix+"%").
		Count(&count).Error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%04d", prefix, count+1), nil
}

func CreateCustomer(c *gin.Context) {
	var req = structs.CustomerCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	customerCode, err := GenerateCustomerCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to generate customer code",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	customer := models.Customer{
		Code:    customerCode,
		Name:    req.Name,
		Phone:   req.Phone,
		Email:   req.Email,
		Npwp:    req.Npwp,
		Address: req.Address,
		Type:    req.Type,
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Customer created successfully",
		Data: structs.CustomerResponse{
			Id:        customer.Id,
			Code:      customer.Code,
			Name:      customer.Name,
			Phone:     customer.Phone,
			Email:     customer.Email,
			Npwp:      customer.Npwp,
			Address:   customer.Address,
			Type:      customer.Type,
			IsActive:  customer.IsActive,
			CreatedAt: customer.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: customer.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindCustomerById(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Customer not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Customer Found",
		Data: structs.CustomerResponse{
			Id:        customer.Id,
			Code:      customer.Code,
			Name:      customer.Name,
			Phone:     customer.Phone,
			Email:     customer.Email,
			Npwp:      customer.Npwp,
			Address:   customer.Address,
			Type:      customer.Type,
			IsActive:  customer.IsActive,
			CreatedAt: customer.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: customer.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Customer not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req structs.CustomerUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	customer.Name = req.Name
	customer.Phone = req.Phone
	customer.Email = req.Email
	customer.Npwp = req.Npwp
	customer.Address = req.Address
	customer.Type = req.Type
	if req.IsActive != nil {
		customer.IsActive = *req.IsActive
	}

	if err := database.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Customer updated successfully",
		Data: structs.CustomerResponse{
			Id:        customer.Id,
			Code:      customer.Code,
			Name:      customer.Name,
			Phone:     customer.Phone,
			Email:     customer.Email,
			Npwp:      customer.Npwp,
			Address:   customer.Address,
			Type:      customer.Type,
			IsActive:  customer.IsActive,
			CreatedAt: customer.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: customer.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Customer not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Customer deleted successfully",
	})
}
