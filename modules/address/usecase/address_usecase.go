package usecase

import (
	"context"
	"golang-contact-management-restful-api/modules/address/models"
)

type AddressUsecase interface {
	Create(ctx context.Context, username string, contactID int, request models.AddressCreateRequest) (models.AddressResponse, error)
	Update(ctx context.Context, username string, contactID int, addressID int, request models.AddressUpdateRequest) (models.AddressResponse, error)
	FindByID(ctx context.Context, username string, contactID int, addressID int) (models.AddressResponse, error)
	FindAll(ctx context.Context, username string, contactID int) ([]models.AddressResponse, error)
	DeleteByID(ctx context.Context, username string, contactID int, addressID int) error
}
