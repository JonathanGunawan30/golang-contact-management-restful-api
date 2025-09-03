package usecase

import (
	"context"
	"errors"
	"golang-contact-management-restful-api/modules/user/domain"
	"golang-contact-management-restful-api/modules/user/entities"
	"golang-contact-management-restful-api/modules/user/models"
	"golang-contact-management-restful-api/modules/user/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecaseImpl struct {
	userRepository repository.UserRepository
	validator      *validator.Validate
}

func NewUserUsecase(userRepository repository.UserRepository, validator *validator.Validate) UserUsecase {
	return &userUsecaseImpl{
		userRepository: userRepository,
		validator:      validator,
	}
}

func (usecase *userUsecaseImpl) Register(ctx context.Context, req models.UserRegisterRequest) (models.UserResponse, error) {
	if err := usecase.validator.Struct(req); err != nil {
		return models.UserResponse{}, err
	}

	isUsernameExists, err := usecase.userRepository.ExistsByUsername(ctx, req.Username)

	if err != nil {
		return models.UserResponse{}, err
	}

	if isUsernameExists {
		return models.UserResponse{}, domain.ErrUsernameAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.UserResponse{}, err
	}

	user := entities.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Name:     req.Name,
	}

	saved, err := usecase.userRepository.Save(ctx, user)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		Username: saved.Username,
		Name:     saved.Name,
	}, nil

}

func (usecase *userUsecaseImpl) Login(ctx context.Context, req models.UserLoginRequest) (models.LoginResponse, error) {
	if err := usecase.validator.Struct(req); err != nil {
		return models.LoginResponse{}, err
	}

	user, err := usecase.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return models.LoginResponse{}, domain.ErrInvalidCredentials
		}
		return models.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.LoginResponse{}, domain.ErrInvalidCredentials
	}

	token := uuid.NewString()

	if _, err = usecase.userRepository.Update(ctx, user.Username, entities.User{Token: token}); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return models.LoginResponse{}, domain.ErrInvalidCredentials
		}
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		Token: token,
	}, nil

}

func (usecase *userUsecaseImpl) UpdateCurrent(ctx context.Context, username string, req models.UserUpdateRequest) (models.UserResponse, error) {
	if err := usecase.validator.Struct(req); err != nil {
		return models.UserResponse{}, err
	}

	var upd entities.User
	if req.Name != "" {
		upd.Name = req.Name
	}

	if req.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.UserResponse{}, err
		}
		upd.Password = string(hashed)
	}

	updated, err := usecase.userRepository.Update(ctx, username, upd)

	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		Username: updated.Username,
		Name:     updated.Name,
	}, nil

}

func (usecase *userUsecaseImpl) GetCurrent(ctx context.Context, username string) (models.UserResponse, error) {
	user, err := usecase.userRepository.FindByUsername(ctx, username)

	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (usecase *userUsecaseImpl) Logout(ctx context.Context, username string) error {
	return usecase.userRepository.ClearTokenByUsername(ctx, username)
}
