package repository

import (
	"context"
	"golang-contact-management-restful-api/modules/address/entities"
)

type AddressRepository interface {
	Save(ctx context.Context, username string, contactID int, address entities.Address) (entities.Address, error)
	UpdateByID(ctx context.Context, username string, contactID int, addressID int, address entities.Address) (entities.Address, error)
	FindByID(ctx context.Context, username string, contactID int, addressID int) (entities.Address, error)
	DeleteByID(ctx context.Context, username string, contactID int, addressID int) error
	FindAll(ctx context.Context, username string, contactID int) ([]entities.Address, error)
}
