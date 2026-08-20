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

func GetAllVendor(c *gin.Context) {
	var vendors []models.Vendor

	database.DB.Find(&vendors)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Get All Vendors",
		Data:    vendors,
	})
}

func GenerateVendorCode() (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("VD-%s-", now.Format("200601"))

	var count int64

	err := database.DB.Model(&models.Vendor{}).
		Where("code LIKE ?", prefix+"%").
		Count(&count).Error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%03d", prefix, count+1), nil
}

func CreateVendor(c *gin.Context) {
	var req = structs.VendorCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	vendorCode, err := GenerateVendorCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to generate vendor code",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	vendor := models.Vendor{
		Code:         vendorCode,
		Name:         req.Name,
		Type:         req.Type,
		Phone:        req.Phone,
		Email:        req.Email,
		Npwp:         req.Npwp,
		Address:      req.Address,
		Note:         req.Note,
		PaymentTerms: req.PaymentTerms,
	}

	if err := database.DB.Create(&vendor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Vendor created successfully",
		Data: structs.VendorResponse{
			Id:           vendor.Id,
			Code:         vendor.Code,
			Name:         vendor.Name,
			Type:         vendor.Type,
			Phone:        vendor.Phone,
			Email:        vendor.Email,
			Npwp:         vendor.Npwp,
			Address:      vendor.Address,
			Note:         vendor.Note,
			PaymentTerms: vendor.PaymentTerms,
			IsActive:     vendor.IsActive,
			CreatedAt:    vendor.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    vendor.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindVendorById(c *gin.Context) {
	id := c.Param("id")

	var vendor models.Vendor

	if err := database.DB.First(&vendor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Vendor not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Vendor found",
		Data: structs.VendorResponse{
			Id:           vendor.Id,
			Code:         vendor.Code,
			Name:         vendor.Name,
			Type:         vendor.Type,
			Phone:        vendor.Phone,
			Email:        vendor.Email,
			Npwp:         vendor.Npwp,
			Address:      vendor.Address,
			Note:         vendor.Note,
			PaymentTerms: vendor.PaymentTerms,
			IsActive:     vendor.IsActive,
			CreatedAt:    vendor.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    vendor.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateVendor(c *gin.Context) {
	id := c.Param("id")

	var vendor models.Vendor

	if err := database.DB.First(&vendor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Vendor not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req structs.VendorUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	vendor.Name = req.Name
	vendor.Type = req.Type
	vendor.Phone = req.Phone
	vendor.Email = req.Email
	vendor.Npwp = req.Npwp
	vendor.Address = req.Address
	vendor.Note = req.Note
	vendor.PaymentTerms = req.PaymentTerms
	if req.IsActive != nil {
		vendor.IsActive = *req.IsActive
	}

	if err := database.DB.Save(&vendor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Vendor updated successfully",
		Data: structs.VendorResponse{
			Id:           vendor.Id,
			Code:         vendor.Code,
			Name:         vendor.Name,
			Type:         vendor.Type,
			Phone:        vendor.Phone,
			Email:        vendor.Email,
			Npwp:         vendor.Npwp,
			Address:      vendor.Address,
			Note:         vendor.Note,
			PaymentTerms: vendor.PaymentTerms,
			IsActive:     vendor.IsActive,
			CreatedAt:    vendor.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    vendor.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteVendor(c *gin.Context) {
	id := c.Param("id")

	var vendor models.Vendor

	if err := database.DB.First(&vendor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Vendor not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&vendor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Vendor deleted successfully",
	})
}
