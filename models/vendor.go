package models

import (
	"time"

	"gorm.io/gorm"
)

type Vendor struct {
	Id           uint           `json:"id" gorm:"primaryKey"`
	Code         string         `json:"code" gorm:"type:varchar(20);unique;not null"`
	Name         string         `json:"name" gorm:"type:varchar(150);not null"`
	Type         string         `json:"type" gorm:"type:enum('SUPPLIER','SUBCON','BOTH'); default:'SUPPLIER'"`
	Phone        string         `json:"phone" gorm:"type:varchar(20)"`
	Email        string         `json:"email" gorm:"type:varchar(150)"`
	Npwp         string         `json:"npwp" gorm:"type:varchar(25)"`
	Address      string         `json:"address" gorm:"type:text"`
	Note         string         `json:"note" gorm:"type:text"`
	PaymentTerms string         `json:"payment_terms" gorm:"type:varchar(3);not null; default:30;"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
