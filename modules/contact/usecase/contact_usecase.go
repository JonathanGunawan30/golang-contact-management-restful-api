package usecase

import (
	"context"
	"golang-contact-management-restful-api/modules/contact/models"
)

type ContactUsecase interface {
	Create(ctx context.Context, username string, request models.ContactCreateRequest) (models.ContactResponse, error)
	Update(ctx context.Context, username string, id int, request models.ContactUpdateRequest) (models.ContactResponse, error)
	FindByID(ctx context.Context, username string, id int) (models.ContactResponse, error)
	DeleteByID(ctx context.Context, username string, id int) error
	Search(ctx context.Context, username string, query models.ContactSearchQuery) ([]models.ContactResponse, models.Paging, error)
}
