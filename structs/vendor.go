package structs

type VendorResponse struct {
	Id           uint   `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Npwp         string `json:"npwp"`
	Address      string `json:"addresss"`
	Note         string `json:"note"`
	PaymentTerms string `json:"payment_terms"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type VendorCreateRequest struct {
	Code         string `json:"code,omitempty" validate:"omitempty"`
	Name         string `json:"name" binding:"required" validate:"required,min=3,max=150"`
	Type         string `json:"type,omitempty" binding:"omitempty,oneof=SUPPLIER SUBCON BOTH" validate:"omitempty,oneof=SUPPLIER SUBCON BOTH"`
	Phone        string `json:"phone,omitempty" validate:"omitempty,max=20"` // Perbaikan typo omitemty
	Email        string `json:"email" binding:"required" validate:"required,email"`
	Npwp         string `json:"npwp,omitempty" validate:"omitempty,max=25"`
	Address      string `json:"address,omitempty" validate:"omitempty"`
	Note         string `json:"note,omitempty" validate:"omitempty"`
	PaymentTerms string `json:"payment_terms" binding:"omitempty"`
	IsActive     *bool  `json:"is_active,omitempty" validate:"omitempty"`
}

type VendorUpdateRequest struct {
	Name         string `json:"name" binding:"required" validate:"required,min=3,max=150"`
	Type         string `json:"type" binding:"required,oneof=SUPPLIER SUBCON BOTH" validate:"omitempty,oneof=SUPPLIER SUBCON BOTH"`
	Phone        string `json:"phone,omitempty" validate:"omitempty,max=20"`
	Email        string `json:"email,omitempty" validate:"omitempty,email"`
	Npwp         string `json:"npwp,omitempty" validate:"omitempty,max=25"`
	Address      string `json:"address,omitempty" validate:"omitempty"`
	Note         string `json:"note,omitempty" validate:"omitempty"`
	PaymentTerms string `json:"payment_terms" binding:"omitempty"`
	IsActive     *bool  `json:"is_active,omitempty" validate:"omitempty"`
}
