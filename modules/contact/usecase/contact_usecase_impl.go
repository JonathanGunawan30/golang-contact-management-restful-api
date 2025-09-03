package usecase

import (
	"context"
	"golang-contact-management-restful-api/internal/transport/http"
	"golang-contact-management-restful-api/modules/contact/entities"
	"golang-contact-management-restful-api/modules/contact/models"
	"golang-contact-management-restful-api/modules/contact/repository"
	"math"

	"github.com/go-playground/validator/v10"
)

type contactUsecaseImpl struct {
	contactRepository repository.ContactRepository
	validator         *validator.Validate
}

func NewContactUsecase(contactRepository repository.ContactRepository, validator *validator.Validate) ContactUsecase {
	return &contactUsecaseImpl{
		contactRepository: contactRepository,
		validator:         validator,
	}
}

func (usecase *contactUsecaseImpl) Create(ctx context.Context, username string, request models.ContactCreateRequest) (models.ContactResponse, error) {
	if err := usecase.validator.Struct(request); err != nil {
		return models.ContactResponse{}, err
	}

	contact := entities.Contact{
		FirstName: request.FirstName,
		LastName:  http.StringToPointerIfNotEmpty(request.LastName),
		Email:     http.StringToPointerIfNotEmpty(request.Email),
		Phone:     http.StringToPointerIfNotEmpty(request.Phone),
	}

	saved, err := usecase.contactRepository.Save(ctx, username, contact)
	if err != nil {
		return models.ContactResponse{}, err
	}

	return models.ContactResponse{
		ID:        saved.ID,
		FirstName: saved.FirstName,
		LastName:  http.PointerToString(saved.LastName),
		Email:     http.PointerToString(saved.Email),
		Phone:     http.PointerToString(saved.Phone),
	}, nil
}

func (usecase *contactUsecaseImpl) Update(ctx context.Context, username string, id int, request models.ContactUpdateRequest) (models.ContactResponse, error) {
	if err := usecase.validator.Struct(request); err != nil {
		return models.ContactResponse{}, err
	}

	contact := entities.Contact{
		FirstName: request.FirstName,
		LastName:  http.StringToPointerIfNotEmpty(request.LastName),
		Email:     http.StringToPointerIfNotEmpty(request.Email),
		Phone:     http.StringToPointerIfNotEmpty(request.Phone),
	}

	updatedContact, err := usecase.contactRepository.UpdateByID(ctx, username, id, contact)
	if err != nil {
		return models.ContactResponse{}, err
	}

	return models.ContactResponse{
		ID:        updatedContact.ID,
		FirstName: updatedContact.FirstName,
		LastName:  http.PointerToString(updatedContact.LastName),
		Email:     http.PointerToString(updatedContact.Email),
		Phone:     http.PointerToString(updatedContact.Phone),
	}, nil
}

func (usecase *contactUsecaseImpl) FindByID(ctx context.Context, username string, id int) (models.ContactResponse, error) {
	retrievedContact, err := usecase.contactRepository.FindByID(ctx, username, id)
	if err != nil {
		return models.ContactResponse{}, err
	}
	return models.ContactResponse{
		ID:        retrievedContact.ID,
		FirstName: retrievedContact.FirstName,
		LastName:  http.PointerToString(retrievedContact.LastName),
		Email:     http.PointerToString(retrievedContact.Email),
		Phone:     http.PointerToString(retrievedContact.Phone),
	}, nil

}

func (usecase *contactUsecaseImpl) DeleteByID(ctx context.Context, username string, id int) error {
	return usecase.contactRepository.DeleteByID(ctx, username, id)
}

func (usecase *contactUsecaseImpl) Search(ctx context.Context, username string, query models.ContactSearchQuery) ([]models.ContactResponse, models.Paging, error) {
	results, total, err := usecase.contactRepository.Search(ctx, username, query)
	if err != nil {
		return nil, models.Paging{}, err
	}

	responses := make([]models.ContactResponse, len(results))
	for i, result := range results {
		responses[i] = models.ContactResponse{
			ID:        result.ID,
			FirstName: result.FirstName,
			LastName:  http.PointerToString(result.LastName),
			Email:     http.PointerToString(result.Email),
			Phone:     http.PointerToString(result.Phone),
		}
	}

	totalPage := int(math.Ceil(float64(total) / float64(query.Size)))
	if totalPage == 0 {
		totalPage = 1
	}

	paging := models.Paging{
		Page:      query.Page,
		TotalPage: totalPage,
		TotalItem: total,
	}

	return responses, paging, nil

}
