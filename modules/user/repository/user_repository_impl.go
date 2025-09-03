package repository

import (
	"context"
	"errors"
	"golang-contact-management-restful-api/modules/user/domain"
	"golang-contact-management-restful-api/modules/user/entities"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repository *userRepositoryImpl) Save(ctx context.Context, user entities.User) (entities.User, error) {
	if err := repository.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return entities.User{}, nil
	}
	return user, nil
}

func (repository *userRepositoryImpl) Update(ctx context.Context, username string, user entities.User) (entities.User, error) {
	updateMap := map[string]any{}
	if user.Name != "" {
		updateMap["name"] = user.Name
	}
	if user.Password != "" {
		updateMap["password"] = user.Password
	}

	if user.Token != "" {
		updateMap["token"] = user.Token
	}

	if len(updateMap) == 0 {
		return entities.User{}, nil
	}

	result := repository.DB.WithContext(ctx).
		Model(&entities.User{}).
		Where("username = ?", username).
		Updates(updateMap)

	if result.Error != nil {
		return entities.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return entities.User{}, domain.ErrUserNotFound
	}

	var updated entities.User
	if err := repository.DB.WithContext(ctx).Take(&updated, "username = ?", username).Error; err != nil {
		return entities.User{}, err
	}

	return updated, nil
}

func (repository *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (entities.User, error) {
	var user entities.User
	if err := repository.DB.WithContext(ctx).Take(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, domain.ErrUserNotFound
		}
		return entities.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	if err := repository.DB.WithContext(ctx).Model(&entities.User{}).
		Where("username = ?", username).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repository *userRepositoryImpl) FindByToken(ctx context.Context, token string) (entities.User, error) {
	var user entities.User
	if err := repository.DB.WithContext(ctx).Take(&user, "token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, domain.ErrUserNotFound
		}
		return entities.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) ClearTokenByUsername(ctx context.Context, username string) error {
	result := repository.DB.WithContext(ctx).Model(&entities.User{}).Where("username = ?", username).Update("token", "")

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
