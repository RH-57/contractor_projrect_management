package models

import "time"

type Customer struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code" gorm:"type:varchar(50);unique;not null"`
	Name      string    `json:"name" gorm:"type:varchar(150);not null;"`
	Type      string    `json:"type" gorm:"type:enum('PRIBADI','PERUSAHAAN','OTHER'); default:'OTHER'"`
	Phone     string    `json:"phone" gorm:"type:varchar(20)"`
	Email     string    `json:"email" gorm:"type:varchar(150)"`
	Npwp      string    `json:"npwp" gorm:"type:varchar(25)"`
	Address   string    `json:"address" gorm:"type:text"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
