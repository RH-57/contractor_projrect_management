package structs

// ProjectResponse digunakan untuk menyajikan data proyek sebagai response API
type ProjectResponse struct {
	Id            uint    `json:"id"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	ClientName    string  `json:"client_name"`
	ClientPhone   string  `json:"client_phone"`
	ContractValue float64 `json:"contract_value"`
	EstimatedCost float64 `json:"estimated_cost"`
	Status        string  `json:"status"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// ProjectCreateRequest digunakan saat membuat proyek baru
type ProjectCreateRequest struct {
	Code          string  `json:"code" binding:"required" validate:"required,min=3,max=50"`
	Name          string  `json:"name" binding:"required" validate:"required,min=3,max=150"`
	ClientName    string  `json:"client_name" binding:"required" validate:"required,min=3,max=100"`
	ClientPhone   string  `json:"client_phone,omitempty" validate:"omitempty,max=20"`
	ContractValue float64 `json:"contract_value" binding:"required" validate:"required,gt=0"`
	EstimatedCost float64 `json:"estimated_cost,omitempty" validate:"omitempty,gte=0"`
	Status        string  `json:"status,omitempty" validate:"omitempty,oneof=PLANNED ON_PROGRESS COMPLETED CANCELLED"`
	StartDate     string  `json:"start_date" binding:"required" validate:"required"` // Format: YYYY-MM-DD
	EndDate       string  `json:"end_date" binding:"required" validate:"required"`   // Format: YYYY-MM-DD
}

// ProjectUpdateRequest digunakan saat memperbarui data proyek
type ProjectUpdateRequest struct {
	Code          string  `json:"code" binding:"required" validate:"required,min=3,max=50"`
	Name          string  `json:"name" binding:"required" validate:"required,min=3,max=150"`
	ClientName    string  `json:"client_name" binding:"required" validate:"required,min=3,max=100"`
	ClientPhone   string  `json:"client_phone,omitempty" validate:"omitempty,max=20"`
	ContractValue float64 `json:"contract_value" binding:"required" validate:"required,gt=0"`
	EstimatedCost float64 `json:"estimated_cost,omitempty" validate:"omitempty,gte=0"`
	Status        string  `json:"status" binding:"required" validate:"required,oneof=PLANNED ON_PROGRESS COMPLETED CANCELLED"`
	StartDate     string  `json:"start_date" binding:"required" validate:"required"` // Format: YYYY-MM-DD
	EndDate       string  `json:"end_date" binding:"required" validate:"required"`   // Format: YYYY-MM-DD
}
