package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

func TranslateErrorMessage(err error) map[string]string {
	errorMessages := make(map[string]string)

	if err == nil {
		return errorMessages
	}

	// 1. Penanganan Error dari Validasi Struct (go-playground/validator)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				errorMessages[field] = fmt.Sprintf("%s is required", e.Field())
			case "min":
				errorMessages[field] = fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
			case "max":
				errorMessages[field] = fmt.Sprintf("%s must not exceed %s characters", e.Field(), e.Param())
			default:
				errorMessages[field] = fmt.Sprintf("%s is invalid", e.Field())
			}
		}
		return errorMessages
	}

	// 2. Penanganan Error dari Database MySQL (misal: Unique Constraint Duplicate Key)
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		// Error code 1062 = Duplicate entry
		if mysqlErr.Number == 1062 {
			if strings.Contains(mysqlErr.Message, "code") {
				errorMessages["code"] = "Code already exists" // Pesan generik yang aman untuk semua tabel
			} else {
				errorMessages["database"] = "Duplicate entry detected"
			}
			return errorMessages
		}
	}

	// 3. Fallback jika error berupa string biasa
	errorMessages["error"] = err.Error()
	return errorMessages
}
