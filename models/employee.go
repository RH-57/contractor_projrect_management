package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Code      string         `json:"code" gorm:"type:varchar(20);unique;not null"`
	Name      string         `json:"name" gorm:"type:varchar(100);not null"`
	Position  string         `json:"position" gorm:"type:varchar(100)"`
	Type      string         `json:"type" gorm:"type:enum('TETAP','KONTRAK','HARIAN','BORONGAN','OTHER'); default:'OTHER'"`
	Phone     string         `json:"phone" gorm:"type:varchar(20)"`
	Address   string         `json:"address" gorm:"type:text"`
	DailyRate float64        `json:"daily_rate" gorm:"type:decimal(15,2);not null; default:0"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
