package usecase

import (
	"context"
	"golang-contact-management-restful-api/modules/user/models"
)

type UserUsecase interface {
	Register(ctx context.Context, req models.UserRegisterRequest) (models.UserResponse, error)
	Login(ctx context.Context, req models.UserLoginRequest) (models.LoginResponse, error)
	UpdateCurrent(ctx context.Context, username string, req models.UserUpdateRequest) (models.UserResponse, error)
	GetCurrent(ctx context.Context, username string) (models.UserResponse, error)
	Logout(ctx context.Context, username string) error
}
