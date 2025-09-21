package models

type SignUpRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type GetNewTokenRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	JwtToken string `json:"jwt_token" validate:"required"`
}

type UpdatePasswordRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	OldPassword string `json:"old_password" validate:"required,min=8,max=100"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=100"`
}