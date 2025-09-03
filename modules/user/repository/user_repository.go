package repository

import (
	"context"
	"golang-contact-management-restful-api/modules/user/entities"
)

type UserRepository interface {
	Save(ctx context.Context, user entities.User) (entities.User, error)
	Update(ctx context.Context, username string, user entities.User) (entities.User, error)
	FindByUsername(ctx context.Context, username string) (entities.User, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	FindByToken(ctx context.Context, token string) (entities.User, error)
	ClearTokenByUsername(ctx context.Context, username string) error
}
