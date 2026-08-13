package structs

// UserResponse digunakan untuk menyajikan data user sebagai response API
type UserResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Role      string  `json:"role"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omitempty"`
}

// UserCreateRequest digunakan saat registrasi / pembuatan user baru
type UserCreateRequest struct {
	Name     string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Username string `json:"username" binding:"required" validate:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,min=6"`
	Role     string `json:"role" binding:"required" validate:"required,oneof=ADMIN PROJECT_MANAGER FINANCE STAFF"`
	IsActive *bool  `json:"is_active,omitempty" validate:"omitempty"`
}

// UserUpdateRequest digunakan saat memperbarui data user
type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Username string `json:"username" binding:"required" validate:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6"`
	Role     string `json:"role,omitempty" validate:"omitempty,oneof=ADMIN PROJECT_MANAGER FINANCE STAFF"`
	IsActive *bool  `json:"is_active,omitempty" validate:"omitempty"`
}

// UserLoginRequest digunakan saat proses otentikasi/login
type UserLoginRequest struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}
