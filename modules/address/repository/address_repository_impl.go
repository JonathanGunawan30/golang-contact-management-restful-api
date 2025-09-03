package repository

import (
	"context"
	"errors"
	domain2 "golang-contact-management-restful-api/modules/address/domain"
	"golang-contact-management-restful-api/modules/address/entities"
	"golang-contact-management-restful-api/modules/contact/domain"
	entities2 "golang-contact-management-restful-api/modules/contact/entities"

	"gorm.io/gorm"
)

type addressRepositoryImpl struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepositoryImpl{DB: db}
}

func (repository *addressRepositoryImpl) Save(ctx context.Context, username string, contactID int, address entities.Address) (entities.Address, error) {
	var contact entities2.Contact
	err := repository.DB.WithContext(ctx).Take(&contact, "username = ? AND id = ?", username, contactID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Address{}, domain.ErrContactNotFound
		}
		return entities.Address{}, err
	}
	address.ContactID = contactID
	if err = repository.DB.WithContext(ctx).Create(&address).Error; err != nil {
		return entities.Address{}, err
	}

	return address, nil

}

func (repository *addressRepositoryImpl) UpdateByID(ctx context.Context, username string, contactID int, addressID int, address entities.Address) (entities.Address, error) {
	var contact entities2.Contact
	err := repository.DB.WithContext(ctx).Where("id= ? AND username = ?", contactID, username).
		Take(&contact).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Address{}, domain.ErrContactNotFound
		}
		return entities.Address{}, err
	}

	updateMap := map[string]any{}

	if address.City != nil {
		updateMap["city"] = *address.City
	}

	if address.Country != "" {
		updateMap["country"] = address.Country
	}

	if address.Street != nil {
		updateMap["street"] = *address.Street
	}

	if address.Province != nil {
		updateMap["province"] = *address.Province
	}

	if address.PostalCode != "" {
		updateMap["postal_code"] = address.PostalCode
	}

	if len(updateMap) == 0 {
		return entities.Address{}, nil
	}

	result := repository.DB.WithContext(ctx).Model(&entities.Address{}).
		Where("id = ? AND contact_id = ? AND contact_id IN (?)", addressID, contactID,
			repository.DB.Model(&entities2.Contact{}).Select("id").Where("username = ?", username),
		).Updates(updateMap)

	if result.Error != nil {
		return entities.Address{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entities.Address{}, domain2.ErrAddressNotFound
	}

	var updated entities.Address
	if err = repository.DB.WithContext(ctx).
		Joins("JOIN contacts ON contacts.id = addresses.contact_id").
		Where("addresses.id = ? AND addresses.contact_id = ? AND contacts.username = ?", addressID, contactID, username).
		Take(&updated).Error; err != nil {
		return entities.Address{}, err
	}

	return updated, nil
}

func (repository *addressRepositoryImpl) FindByID(ctx context.Context, username string, contactID int, addressID int) (entities.Address, error) {
	var address entities.Address
	if err := repository.DB.WithContext(ctx).Model(&entities.Address{}).Joins("JOIN contacts ON contacts.id = addresses.contact_id").
		Where("addresses.id = ? AND addresses.contact_id = ? AND contacts.username = ?", addressID, contactID, username).
		Take(&address).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Address{}, domain2.ErrAddressNotFound
		}
		return entities.Address{}, err
	}
	return address, nil
}

func (repository *addressRepositoryImpl) DeleteByID(ctx context.Context, username string, contactID int, addressID int) error {
	var address entities.Address
	err := repository.DB.WithContext(ctx).
		Joins("JOIN contacts ON contacts.id = addresses.contact_id").
		Where("addresses.id = ? AND addresses.contact_id = ? AND contacts.username = ?", addressID, contactID, username).
		First(&address).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain2.ErrAddressNotFound
		}
		return err
	}

	result := repository.DB.WithContext(ctx).Delete(&entities.Address{}, address.ID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain2.ErrAddressNotFound
	}

	return nil
}

func (repository *addressRepositoryImpl) FindAll(ctx context.Context, username string, contactID int) ([]entities.Address, error) {
	var addresses []entities.Address
	if err := repository.DB.WithContext(ctx).Model(&entities.Address{}).Joins("JOIN contacts ON contacts.id = addresses.contact_id").
		Where("contacts.id = ? AND contacts.username = ?", contactID, username).Find(&addresses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrContactNotFound
		}
		return nil, err
	}

	return addresses, nil

}
