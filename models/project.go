package models

import "time"

type Project struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	Code          string    `json:"code" gorm:"type:varchar(50);unique;not null"`
	Name          string    `json:"name" gorm:"type:varchar(150);not null"`
	CustomerID    uint      `json:"customer_id" gorm:"not null"`
	Customer      Customer  `json:"customer" gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ContractValue float64   `json:"contract_value" gorm:"decimal(15,2);not null;default:0"`
	EstimatedCost float64   `json:"estimated_cost" gorm:"type:decimal(15,2);default:0"`
	Status        string    `json:"status" gorm:"enum:('PLANNED','ON_PROGRESS','COMPLETE','CANCELED');default:'PLANNED'"`
	StartDate     time.Time `json:"start_date" gorm:"type:date"`
	EndDate       time.Time `json:"end_date" gorm:"type:date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
