package repository

import (
	"context"
	"golang-contact-management-restful-api/modules/contact/entities"
	"golang-contact-management-restful-api/modules/contact/models"
)

type ContactRepository interface {
	Save(ctx context.Context, username string, contact entities.Contact) (entities.Contact, error)
	UpdateByID(ctx context.Context, username string, id int, contact entities.Contact) (entities.Contact, error)
	FindByID(ctx context.Context, username string, id int) (entities.Contact, error)
	DeleteByID(ctx context.Context, username string, id int) error
	Search(ctx context.Context, username string, query models.ContactSearchQuery) ([]entities.Contact, int, error)
}
