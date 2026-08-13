package database

import (
	"log"
	"projectmanagement/backend-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB) {
	var count int64

	db.Model(&models.User{}).Count(&count)

	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin12345"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password for seeder: %v", err)
		}

		admin := models.User{
			Name:     "Super Admin",
			Username: "admin",
			Email:    "admin@kspkontraktor.com",
			Password: string(hashedPassword),
			Role:     "ADMIN",
			IsActive: true,
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Printf("Failed to seed admin user: %v", err)
		} else {
			log.Println("=== First Admin User Created Successfully ===")
			log.Println("Username: admin | Password: admin123")
		}
	}
}
