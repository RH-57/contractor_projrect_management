package models

import "time"

type Setting struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	AppName     string    `json:"app_name" gorm:"type:varchar(100);default:'KSP Kontraktor'"`
	CompanyName string    `json:"company_name" gorm:"type:varchar(150)"`
	Phone       string    `json:"phone" gorm:"type:varchar(30)"`
	Email       string    `json:"email" gorm:"type:varchar(150)"`
	Npwp        string    `json:"npwp" gorm:"type:varchar(30)"`
	Address     string    `json:"address" gorm:"type:text"`
	Logo        string    `json:"logo" gorm:"type:varchar(255)"`                   // URL / Path file logo
	TaxRate     float64   `json:"tax_rate" gorm:"type:decimal(5,2);default:11.00"` // e.g., 11.00 (%)
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
