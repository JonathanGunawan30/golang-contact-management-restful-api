package usecase

import (
	"context"
	"golang-contact-management-restful-api/internal/transport/http"
	"golang-contact-management-restful-api/modules/address/entities"
	"golang-contact-management-restful-api/modules/address/models"
	"golang-contact-management-restful-api/modules/address/repository"

	"github.com/go-playground/validator/v10"
)

type addressUsecaseImpl struct {
	addressRepository repository.AddressRepository
	validator         *validator.Validate
}

func NewAddressUsecase(addressRepository repository.AddressRepository, validator *validator.Validate) AddressUsecase {
	return &addressUsecaseImpl{
		addressRepository: addressRepository,
		validator:         validator,
	}
}

func (usecase *addressUsecaseImpl) Create(ctx context.Context, username string, contactID int, request models.AddressCreateRequest) (models.AddressResponse, error) {
	if err := usecase.validator.Struct(request); err != nil {
		return models.AddressResponse{}, err
	}

	address := entities.Address{
		Street:     http.StringToPointerIfNotEmpty(request.Street),
		City:       http.StringToPointerIfNotEmpty(request.City),
		Province:   http.StringToPointerIfNotEmpty(request.Province),
		Country:    request.Country,
		PostalCode: request.PostalCode,
	}

	saved, err := usecase.addressRepository.Save(ctx, username, contactID, address)
	if err != nil {
		return models.AddressResponse{}, err
	}

	return models.AddressResponse{
		ID:         saved.ID,
		Street:     http.PointerToString(saved.Street),
		City:       http.PointerToString(saved.City),
		Province:   http.PointerToString(saved.Province),
		Country:    saved.Country,
		PostalCode: saved.PostalCode,
	}, nil

}

func (usecase *addressUsecaseImpl) Update(ctx context.Context, username string, contactID int, addressID int, request models.AddressUpdateRequest) (models.AddressResponse, error) {
	if err := usecase.validator.Struct(request); err != nil {
		return models.AddressResponse{}, err
	}

	address := entities.Address{
		Street:     http.StringToPointerIfNotEmpty(request.Street),
		City:       http.StringToPointerIfNotEmpty(request.City),
		Province:   http.StringToPointerIfNotEmpty(request.Province),
		Country:    request.Country,
		PostalCode: request.PostalCode,
	}

	updatedAddress, err := usecase.addressRepository.UpdateByID(ctx, username, contactID, addressID, address)
	if err != nil {
		return models.AddressResponse{}, err
	}

	return models.AddressResponse{
		ID:         updatedAddress.ID,
		Street:     http.PointerToString(updatedAddress.Street),
		City:       http.PointerToString(updatedAddress.City),
		Province:   http.PointerToString(updatedAddress.Province),
		Country:    updatedAddress.Country,
		PostalCode: updatedAddress.PostalCode,
	}, nil

}

func (usecase *addressUsecaseImpl) FindByID(ctx context.Context, username string, contactID int, addressID int) (models.AddressResponse, error) {
	result, err := usecase.addressRepository.FindByID(ctx, username, contactID, addressID)
	if err != nil {
		return models.AddressResponse{}, err
	}

	return models.AddressResponse{
		ID:         result.ID,
		Street:     http.PointerToString(result.Street),
		City:       http.PointerToString(result.City),
		Province:   http.PointerToString(result.Province),
		Country:    result.Country,
		PostalCode: result.PostalCode,
	}, nil
}

func (usecase *addressUsecaseImpl) FindAll(ctx context.Context, username string, contactID int) ([]models.AddressResponse, error) {
	results, err := usecase.addressRepository.FindAll(ctx, username, contactID)
	if err != nil {
		return nil, err
	}

	responses := make([]models.AddressResponse, len(results))
	for i, result := range results {
		responses[i] = models.AddressResponse{
			ID:         result.ID,
			Street:     http.PointerToString(result.Street),
			City:       http.PointerToString(result.City),
			Province:   http.PointerToString(result.Province),
			Country:    result.Country,
			PostalCode: result.PostalCode,
		}
	}

	return responses, nil
}

func (usecase *addressUsecaseImpl) DeleteByID(ctx context.Context, username string, contactID int, addressID int) error {
	return usecase.addressRepository.DeleteByID(ctx, username, contactID, addressID)
}
