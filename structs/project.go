package structs

type ProjectResponse struct {
	Id            uint              `json:"id"`
	Code          string            `json:"code"`
	Name          string            `json:"name"`
	CustomerID    uint              `json:"customer_id"`
	Customer      *CustomerResponse `json:"customer,omitempty"`
	ContractValue float64           `json:"contract_value"`
	EstimatedCost float64           `json:"estimated_cost"`
	Status        string            `json:"status"`
	StartDate     string            `json:"start_date"`
	EndDate       string            `json:"end_date"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
}

type ProjectCreateRequest struct {
	Code          string  `json:"code,omitempty" validate:"omitempty"`
	Name          string  `json:"name" binding:"required" validate:"required,min=3,max=150"`
	CustomerID    uint    `json:"customer_id" binding:"required" validate:"required,gt=0"`
	ContractValue float64 `json:"contract_value" binding:"required" validate:"required,gt=0"`
	EstimatedCost float64 `json:"estimated_cost,omitempty" validate:"omitempty,gte=0"`
	Status        string  `json:"status,omitempty" validate:"omitempty,oneof=PLANNED ON_PROGRESS COMPLETED CANCELLED"`
	StartDate     string  `json:"start_date" binding:"required" validate:"required"`
	EndDate       string  `json:"end_date" binding:"required" validate:"required"`
}

type ProjectUpdateRequest struct {
	Name          string  `json:"name" binding:"required" validate:"required,min=3,max=150"`
	CustomerID    uint    `json:"customer_id" binding:"required" validate:"required,gt=0"`
	ContractValue float64 `json:"contract_value" binding:"required" validate:"required,gt=0"`
	EstimatedCost float64 `json:"estimated_cost,omitempty" validate:"omitempty,gte=0"`
	Status        string  `json:"status" binding:"required" validate:"required,oneof=PLANNED ON_PROGRESS COMPLETED CANCELLED"`
	StartDate     string  `json:"start_date" binding:"required" validate:"required"`
	EndDate       string  `json:"end_date" binding:"required" validate:"required"`
}
