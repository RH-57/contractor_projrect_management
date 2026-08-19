package structs

type EmployeeResponse struct {
	Id        uint    `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Position  string  `json:"position"`
	Type      string  `json:"type"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	DailyRate float64 `json:"daily_rate"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

type EmployeeCreateRequest struct {
	Code      string  `json:"code,omitempty" validate:"omitempty"`
	Name      string  `json:"name" binding:"required" validate:"required,min=3,max=150"`
	Position  string  `json:"position" binding:"required" validate:"required,max=100"`
	Type      string  `json:"type,omitempty" binding:"omitempty,oneof=TETAP KONTRAK HARIAN BORONGAN OTHER" validate:"omitempty,oneof=TETAP KONTRAK HARIAN BORONGAN OTHER"`
	Phone     string  `json:"phone,omitempty" validate:"omitempty,max=20"`
	Address   string  `json:"address,omitempty" validate:"omitempty"`
	DailyRate float64 `json:"daily_rate" binding:"required" validate:"required,gte=0"`
	IsActive  *bool   `json:"is_active,omitempty" validate:"omitempty"`
}

type EmployeeUpdateRequest struct {
	Name      string  `json:"name" binding:"required" validate:"required,min=3,max=150"`
	Position  string  `json:"position" binding:"required" validate:"required,max=100"`
	Type      string  `json:"type" binding:"required,oneof=TETAP KONTRAK HARIAN BORONGAN OTHER" validate:"required,oneof=TETAP KONTRAK HARIAN BORONGAN OTHER"`
	Phone     string  `json:"phone,omitempty" validate:"omitempty,max=20"`
	Address   string  `json:"address,omitempty" validate:"omitempty"`
	DailyRate float64 `json:"daily_rate" binding:"required" validate:"required,gte=0"`
	IsActive  *bool   `json:"is_active,omitempty" validate:"omitempty"`
}
