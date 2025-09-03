package models

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Name     string `json:"name"     validate:"required,min=3,max=100"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type UserUpdateRequest struct {
	Name     string `json:"name"     validate:"omitempty,min=3,max=100"`
	Password string `json:"password" validate:"omitempty,min=8,max=255"`
}
type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
