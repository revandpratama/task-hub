package dto

type LoginRequest struct {
	Identifier string `json:"identifier" form:"identifier" validate:"required,min=3,max=32"`
	Password   string `json:"password" form:"password" validate:"required,min=8,max=128"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min:4"`
	Username string `json:"username" form:"username" validate:"required,min=4,max=32,alphanum"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=128"`
}
