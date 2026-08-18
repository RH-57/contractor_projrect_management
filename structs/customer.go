package structs

type CustomerResponse struct {
	Id        uint   `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Npwp      string `json:"npwp"`
	Address   string `json:"address"`
	Type      string `json:"type"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CustomerCreateRequest struct {
	Code     string `json:"code,omitempty" validate:"omitempty"`
	Name     string `json:"name" binding:"required" validate:"required,min=3,max=150"`
	Phone    string `json:"phone,omitempty" validate:"omitempty,max=20"` // Perbaikan typo omitemty
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Npwp     string `json:"npwp,omitempty" validate:"omitempty,max=25"`
	Address  string `json:"address,omitempty" validate:"omitempty"`
	Type     string `json:"type,omitempty" binding:"omitempty,oneof=PRIBADI PERUSAHAAN OTHER" validate:"omitempty,oneof=PRIBADI PERUSAHAAN OTHER"`
	IsActive *bool  `json:"is_active,omitempty" validate:"omitempty"`
}

type CustomerUpdateRequest struct {
	Name     string `json:"name" binding:"required" validate:"required,min=3,max=150"`
	Type     string `json:"type" binding:"required,oneof=PRIBADI PERUSAHAAN OTHER" validate:"required,oneof=PRIBADI PERUSAHAAN OTHER"`
	Phone    string `json:"phone,omitempty" validate:"omitempty,max=20"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Npwp     string `json:"npwp,omitempty" validate:"omitempty,max=25"`
	Address  string `json:"address,omitempty" validate:"omitempty"`
	IsActive *bool  `json:"is_active,omitempty" validate:"omitempty"`
}
