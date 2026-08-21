package structs

// UpdateSettingRequest digunakan untuk menangani payload HTTP PUT/PATCH
type UpdateSettingRequest struct {
	AppName     string  `json:"app_name" binding:"required"`
	CompanyName string  `json:"company_name"`
	Phone       string  `json:"phone" binding:"required"`
	Email       string  `json:"email" binding:"required,email"`
	Npwp        string  `json:"npwp"`
	Address     string  `json:"address"`
	Logo        string  `json:"logo"`
	TaxRate     float64 `json:"tax_rate" binding:"gte=0,lte=100"` // Validasi persentase 0-100%
}

// SettingResponse digunakan untuk response JSON ke frontend
type SettingResponse struct {
	Id          uint    `json:"id"`
	AppName     string  `json:"app_name"`
	CompanyName string  `json:"company_name"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	Npwp        string  `json:"npwp"`
	Address     string  `json:"address"`
	Logo        string  `json:"logo"`
	TaxRate     float64 `json:"tax_rate"`
	UpdatedAt   string  `json:"updated_at"`
}
