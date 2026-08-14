package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	errorsMap := make(map[string]string)

	if err == nil {
		return errorsMap
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()

			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorsMap[field] = "Invalid email format"
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field)
			case "oneof":
				errorsMap[field] = fmt.Sprintf("%s must be one of %s", field, fieldError.Param())
			default:
				errorsMap[field] = fmt.Sprintf("%s is invalid", field)
			}
		}

		return errorsMap
	}

	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		errMsg := mysqlErr.Message

		if strings.Contains(errMsg, "username") || strings.Contains(errMsg, "idx_users_username") {
			errorsMap["Username"] = "Username already exists"
		} else if strings.Contains(errMsg, "email") || strings.Contains(errMsg, "idx_users_email") {
			errorsMap["Email"] = "Email already exists"
		} else if strings.Contains(errMsg, "code") || strings.Contains(errMsg, "idx_projects_code") {
			errorsMap["Code"] = "Project code already exists"
		} else {
			errorsMap["Database"] = "Duplicate entry error"
		}
		return errorsMap
	}

	if strings.Contains(err.Error(), "Duplicate entry") {
		if strings.Contains(err.Error(), "username") {
			errorsMap["Username"] = "Username already exists"
		} else if strings.Contains(err.Error(), "email") {
			errorsMap["Email"] = "Email already exists"
		} else if strings.Contains(err.Error(), "code") {
			errorsMap["Code"] = "Project code already exists"
		} else {
			errorsMap["Database"] = "Duplicate entry error"
		}
		return errorsMap
	}

	if err == gorm.ErrRecordNotFound {
		errorsMap["Error"] = "Record not found"
		return errorsMap
	}

	// 4. Fallback Error Lainnya
	errorsMap["Error"] = err.Error()
	return errorsMap
}

func IsDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}

	// Cek langsung dari tipe MySQL Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return mysqlErr.Number == 1062
	}

	// Cek berdasarkan substring pesan error
	return strings.Contains(err.Error(), "Duplicate entry")
}
